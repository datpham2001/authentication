package repository

import (
	"authentication/dto"
	"authentication/model"
)

type AuthRepository interface {
	SignUp(input *dto.AccountRegisterRequest) (*model.Account, dto.DatabaseError)
	Login(input *dto.AccountLoginRequest) (*model.Account, dto.DatabaseError)
	Logout(input *dto.AccountLoginRequest) (*model.Account, dto.DatabaseError)
}
