package main

import "os"

// ImageGetter fetch image
type ImageGetter interface {
	GetImage() (name string, err error)
}

// Informator provide information about images path
type Informator interface {
	GetInfo() []os.FileInfo
}
