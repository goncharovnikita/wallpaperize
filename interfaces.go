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

// WallMaster is responsible for setting up os wallpapers
type WallMaster interface {
	Get() (string, error)
	SetFromFile(file string) error
}
