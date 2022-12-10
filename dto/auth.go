package dto

type AccountRegisterRequest struct {
	Account struct {
		Email           string `json:"email,omitempty" bson:"email,omitempty"`
		Username        string `json:"username,omitempty" bson:"username,omitempty"`
		Password        string `json:"password,omitempty" bson:"password,omitempty"`
		PasswordConfirm string `json:"passwordConfirm,omitempty" bson:"password_confirm,omitempty"`
	} `json:"account"`
}

type AccountLoginRequest struct {
	Account struct {
		Email    string `json:"email,omitempty" bson:"email,omitempty"`
		Password string `json:"password,omitempty" bson:"password,omitempty"`
	} `json:"account"`
}
