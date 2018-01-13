package main

import "io"

// RandomImageGetter interface provide get random inage method
type RandomImageGetter interface {
	GetRandomImage() ([]byte, error)
}

// RandomImageReader interface
type RandomImageReader interface {
	GetRandomImageReader() (io.ReadCloser, error)
}

// ImageReaderGetter interface
type ImageReaderGetter interface {
	GetImageReader() (io.ReadCloser, error)
}
