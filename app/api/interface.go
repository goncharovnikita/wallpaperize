package api

// Wallmaster get and set wallpaper from filepath
type Wallmaster interface {
	Get() (string, error)
	SetFromFile(file string) error
}

// ImageGetter fetch image
type ImageGetter interface {
	GetImage() (img []byte, err error)
}
