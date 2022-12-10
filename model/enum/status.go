package enum

type StatusValue struct {
	Ok           string
	Error        string
	Invalid      string
	NotFound     string
	Forbidden    string
	Existed      string
	Unauthorized string
}

var APIStatus = &StatusValue{
	Ok:           "OK",
	Error:        "ERROR",
	Invalid:      "INVALID",
	NotFound:     "NOT_FOUND",
	Forbidden:    "FORBIDDEN",
	Existed:      "EXISTED",
	Unauthorized: "UNAUTHORIZED",
}
