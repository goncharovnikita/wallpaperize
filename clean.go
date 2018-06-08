package main

import (
	"os"
	"path/filepath"
)

// cleaner type
type cleaner struct{}

// clean all outdated random photos
func (c cleaner) cleanRandomImages() (err error) {
	var names configStructure
	if names, err = conf.parseConfig(); err != nil {
		return
	}

	return filepath.Walk(absRandomDirname, func(path string, info os.FileInfo, e error) error {
		exists := false
		for _, v := range names.RandomPhotos {
			if v.Name == info.Name() || info.Name() == "config" {
				exists = true
				break
			}
		}
		if !exists {
			os.Remove(absRandomDirname + "/" + info.Name())
		}
		return nil
	})
}
