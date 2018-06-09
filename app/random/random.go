package random

import (
	"bytes"
	"io"
	"log"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/hash"
)

// Random image getter type
type Random struct {
	rg   RandomImageGetter
	path string
}

// NewRandomImageGetter create new random image getter
func NewRandomImageGetter(rg RandomImageGetter, path string) *Random {
	return &Random{
		rg:   rg,
		path: path,
	}
}

// GetImage implementation
func (r *Random) GetImage() (string, error) {
	img, err := r.rg.GetRandomImage()
	if err != nil {
		log.Println(err)
		return "", err
	}

	hsh := hash.Hash256(img)
	fname := r.path + "/random" + hsh + ".jpg"

	file, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println(err)
		return "", err
	}

	defer file.Close()

	_, err = io.Copy(file, bytes.NewReader(img))
	if err != nil {
		return "", err
	}

	return fname, nil
}
