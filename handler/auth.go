package handler

import (
	"authentication/dto"
	"authentication/model/common"
	"authentication/model/enum"
	"authentication/pkg"
	"authentication/pkg/utils"
	"authentication/repository"
	"net/http"

	"github.com/labstack/echo"
)

type authHandler struct {
	Repository repository.AuthRepository
}

func NewAuthHandler(repo repository.AuthRepository) *authHandler {
	return &authHandler{
		Repository: repo,
	}
}

func (h *authHandler) SignUp(c echo.Context) error {
	var input dto.AccountRegisterRequest

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Parse request data error. " + err.Error(),
			ErrorCode: string(enum.ErrorCodeInvalid.ParseData),
		})
	}

	// validate input
	account := input.Account
	if account.Email == "" {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Email is required",
			ErrorCode: string(enum.ErrorCodeRequired.Email),
		})
	}

	if account.Username == "" {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Username is required",
			ErrorCode: string(enum.ErrorCodeRequired.Username),
		})
	}

	if account.Password == "" {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Password is required",
			ErrorCode: string(enum.ErrorCodeRequired.Password),
		})
	}

	if account.Email == "" {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Password confirm is required",
			ErrorCode: string(enum.ErrorCodeRequired.PasswordConfirm),
		})
	}

	// validate regex
	if !utils.ValidateEmail(account.Email) {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Email is invalid",
			ErrorCode: string(enum.ErrorCodeInvalid.Email),
		})
	}

	if !utils.ValidateUsername(account.Username) {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Username is invalid",
			ErrorCode: string(enum.ErrorCodeInvalid.Username),
		})
	}

	if !utils.ValidatePassword(account.Password) {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Password is invalid",
			ErrorCode: string(enum.ErrorCodeInvalid.Password),
		})
	}

	if account.PasswordConfirm != account.Password {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Password confirm must be matched with password",
			ErrorCode: string(enum.ErrorCodeInvalid.PasswordConfirm),
		})
	}

	accountResp, dbErr := h.Repository.SignUp(&input)
	if dbErr.Code != enum.ErrorCodeValue("SUCCESS") {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   pkg.GetMessageByCode(dbErr.Code),
			ErrorCode: string(enum.ErrorDatabaseCode.Create),
		})
	}

	return c.JSON(http.StatusCreated, &dto.APIResponse{
		Status:  common.APIStatus.Ok,
		Message: "Create new account successfully",
		Data:    accountResp,
	})
}

func (h *authHandler) Login(c echo.Context) error {
	var input dto.AccountLoginRequest

	err := c.Bind(&input)
	if err != nil {
		return c.JSON(http.StatusBadRequest, &dto.APIResponse{
			Status:    common.APIStatus.Invalid,
			Message:   "Parse request data error. " + err.Error(),
			ErrorCode: string(enum.ErrorCodeInvalid.ParseData),
		})
	}

}
