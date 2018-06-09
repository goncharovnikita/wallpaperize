// Package daily encapsulates work with daily images
package daily

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goncharovnikita/wallpaperize/app/cerrors"
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
	return d.getImage()
}

func (d *Daily) getImage() (name string, err error) {
	name = d.path + "/" + getToday() + ".jpg"
	info, err := os.Stat(name)
	if err != nil {
		if err.Error() == cerrors.NewStatNoSuchFileErr(name).Error() {
			img, err := d.dg.GetDailyImage()
			if err != nil {
				log.Println(err)
				return "", err
			}

			file, err := os.OpenFile(name, os.O_CREATE|os.O_RDWR, 0777)
			if err != nil {
				log.Println(err)
				return "", err
			}

			defer file.Close()

			_, err = file.Write(img)
			if err != nil {
				log.Println(err)
				return "", err
			}
			return name, nil
		}
		log.Println(err)
		return "", err
	}

	if info.IsDir() {
		return "", fmt.Errorf("Path is directory, please delete and retry again - %s", name)
	}

	return name, nil
}

func getToday() string {
	now := time.Now()
	year, month, day := now.Date()
	return fmt.Sprintf("%d.%s.%d", day, month, year)
}
