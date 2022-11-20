package api

import "authentication/dto"

type RegisterRequest struct {
	User dto.AccountRegisterRequest `json:"user"`
}
