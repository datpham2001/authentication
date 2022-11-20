package dao

import "authentication/model"

type AuthDAO interface {
	Register(input *model.Account) (*model.Account, error)
	Login(input *model.Account) (*model.Account, error)
}
