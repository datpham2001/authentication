package enum

type ErrorCodeValue string

// INVALID CODE
type errorCodeInvalid struct {
	ParseData       ErrorCodeValue
	Email           ErrorCodeValue
	Password        ErrorCodeValue
	Username        ErrorCodeValue
	PasswordConfirm ErrorCodeValue
}

var ErrorCodeInvalid = &errorCodeInvalid{
	ParseData:       "INVALID_PARSE_DATA",
	Email:           "INVALID_EMAIL",
	Password:        "INVALID_PASSWORD",
	Username:        "INVALID_USERNAME",
	PasswordConfirm: "INVALID_PASSWORD_CONFIRM",
}

// EXISTED ERROR CODE
type errorCodeExisted struct {
	Account ErrorCodeValue
}

var ErrorCodeExisted = &errorCodeExisted{
	Account: "EXISTED_ACCOUNT",
}

type errorCodeRequired struct {
	Email           ErrorCodeValue
	Username        ErrorCodeValue
	Password        ErrorCodeValue
	PasswordConfirm ErrorCodeValue
}

var ErrorCodeRequired = &errorCodeRequired{
	Email:           "REQUIRED_EMAIL",
	Username:        "REQUIRED_USERNAME",
	Password:        "REQUIRED_PASSWORD",
	PasswordConfirm: "REQUIRED_PASSWORD_CONFIRM",
}

type errorDatabaseCode struct {
	Create ErrorCodeValue
	Get    ErrorCodeValue
	Update ErrorCodeValue
	Delete ErrorCodeValue
}

var ErrorDatabaseCode = &errorDatabaseCode{
	Create: "ERROR_CREATE_DOCUMENT",
	Get:    "ERROR_GET_DOCUMENT",
	Update: "ERROR_UPDATE_DOCUMENT",
	Delete: "ERROR_DELETE_DOCUMENT",
}
