// +build windows

package wallmaster

import "github.com/reujab/wallpaper"

type Wallmaster struct{}

func (Wallmaster) Get() (string, error) {
	return wallpaper.Get()
}

func (Wallmaster) SetFromFile(s string) error {
	return wallpaper.SetFromFile(s)
}
