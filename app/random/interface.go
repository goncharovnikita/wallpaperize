package random

// RandomImageGetter fetch random image
type RandomImageGetter interface {
	GetRandomImage() ([]byte, error)
}
