package main

import (
	"log"
	"os/exec"
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

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}

}
