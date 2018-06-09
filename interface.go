package main

// ImageGetter fetch image
type ImageGetter interface {
	GetImage() ([]byte, error)
}
