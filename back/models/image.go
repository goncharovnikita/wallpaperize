package models

type DBImage struct {
	Data []byte
}

type ResponseImage struct {
	Data []byte `json:"data"`
}

func MakeResponseImage(image *DBImage) *ResponseImage {
	return &ResponseImage{
		Data: image.Data,
	}
}

func MakeDBImage(image *ResponseImage) *DBImage {
	return &DBImage{
		Data: image.Data,
	}
}

type ResponseError struct {
	Error string `json:"error"`
}
