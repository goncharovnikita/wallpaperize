// Package daily encapsulates work with daily images
package daily

import (
	"errors"
	"fmt"
	"os"
	"time"
)

// Daily fetch daily images
type Daily struct {
	dg   DailyGetter
	path string
}

// NewDailyGetter creates new daily getter instance
func NewDailyGetter(dg DailyGetter, cachePath string) *Daily {
	return &Daily{
		dg:   dg,
		path: cachePath,
	}
}

// GetImage ImageGetter implementation
func (d *Daily) GetImage() (string, error) {
	name := d.path + "/" + getToday() + ".jpg"

	info, err := os.Stat(name)
	if err != nil {
		if errors.Is(err, &os.PathError{}) {
			img, err := d.dg.GetDailyImage()
			if err != nil {
				return "", err
			}

			file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				return "", err
			}

			defer file.Close()

			_, err = file.Write(img)
			if err != nil {
				return "", err
			}

			return name, nil
		}

		return "", err
	}

	if info.IsDir() {
		return "", fmt.Errorf("path is directory, please delete and retry again - %s", name)
	}

	return name, nil
}

func getToday() string {
	year, month, day := time.Now().Date()

	return fmt.Sprintf("%d.%s.%d", day, month, year)
}
