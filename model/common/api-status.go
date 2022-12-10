package common

type apiStatusValue struct {
	Ok        string
	Notfound  string
	Invalid   string
	Forbidden string
}

var APIStatus = &apiStatusValue{
	Ok:        "OK",
	Notfound:  "Notfound",
	Invalid:   "Invalid",
	Forbidden: "Forbidden",
}
