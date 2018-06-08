package main

import (
	"log"
	"path/filepath"
)

// set's photo of the day as wallpaper
func setPhotoOfTheDay() {
	var (
		filename string
		absPath  string
		err      error
	)

	if filename, err = ch.retrieve(false); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs(absCacheDirname + filename); err != nil {
		log.Fatal(err)
	}

	if err = wallmaster.SetFromFile(absPath); err != nil {
		log.Fatal(err)
	}

}
