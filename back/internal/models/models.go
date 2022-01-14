package models

type ResponseError struct {
	Error string `json:"error"`
}

type DBImage struct {
	Data []byte
}
