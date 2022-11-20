package dto

type AccountRegisterRequest struct {
	Email           string `json:"email,omitempty" bson:"email,omitempty"`
	Username        string `json:"username,omitempty" bson:"username,omitempty"`
	Password        string `json:"password,omitempty" bson:"password,omitempty"`
	PasswordConfirm string `json:"passwordConfirm,omitempty" bson:"password_confirm,omitempty"`
}

type AccountLoginRequest struct {
	Email    string `json:"email,omitempty" bson:"email,omitempty"`
	Password string `json:"password,omitempty" bson:"password,omitempty"`
}
