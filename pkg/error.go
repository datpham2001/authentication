package pkg

import "authentication/model/enum"

var DatabaseErrorMessageMap map[enum.ErrorCodeValue]string = map[enum.ErrorCodeValue]string{
	enum.ErrorCodeExisted.Account: "Username/Email is existed",
	enum.ErrorDatabaseCode.Create: "Error when create a new document",
	enum.ErrorDatabaseCode.Get:    "Error when get document(s)",
	enum.ErrorDatabaseCode.Update: "Error when update document(s)",
	enum.ErrorDatabaseCode.Delete: "Error when delete document(s)",
}

func GetMessageByCode(code enum.ErrorCodeValue) string {
	return DatabaseErrorMessageMap[code]
}
