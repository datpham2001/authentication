package dto

import "authentication/model/enum"

type APIResponse struct {
	Status    string
	Message   string
	ErrorCode string
	Data      interface{}
}

type DatabaseError struct {
	Code enum.ErrorCodeValue
}
