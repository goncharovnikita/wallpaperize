package main

import (
	"log"
	"os/exec"
	"path/filepath"
)

func setRandomPhoto() {
	var (
		absPath  string
		err      error
		filename string
	)

	if filename, err = ch.retrieve(true); err != nil {
		log.Fatal(err)
	}

	if absPath, err = filepath.Abs(absRandomDirname + "/" + filename); err != nil {
		log.Fatal(err)
	}

	if err = exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", absPath).Run(); err != nil {
		log.Fatal(err)
	}
}
