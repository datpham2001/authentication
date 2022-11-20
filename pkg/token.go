package pkg

import (
	"authentication/config"
	"authentication/model/enum"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateTokenPairs(accountID string, accountRole enum.AccountRoleEnumValue) (map[string]string, error) {
	// create access token
	accessToken := jwt.New(jwt.SigningMethodHS256)

	// set claims
	claims := accessToken.Claims.(jwt.MapClaims)
	claims["sub"] = accountID
	claims["role"] = accountRole
	claims["exp"] = time.Now().Add(5 * time.Minute).Unix()

	// generate access token and send it as response
	accToken, err := accessToken.SignedString(config.AppConfig.JWTSecret)
	if err != nil {
		return nil, err
	}

	// same with refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)

	exClaims := refreshToken.Claims.(jwt.MapClaims)
	exClaims["sub"] = accountID
	exClaims["exp"] = time.Now().Add(7 * 24 * time.Hour).Unix()

	refToken, err := refreshToken.SignedString(config.App)
	if err != nil {
		return nil, err
	}

	return map[string]string{
		"access_token":  accToken,
		"refresh_token": refToken,
	}, nil
}
