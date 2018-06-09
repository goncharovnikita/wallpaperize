package main

import (
	"fmt"
	"log"
	"os"

	"github.com/goncharovnikita/wallpaperize/app/cerrors"
)

func ensureDir(path string) {
	info, err := os.Stat(path)
	if err != nil {
		if err.Error() == cerrors.NewStatNoSuchFileErr(path).Error() {
			err = os.Mkdir(path, 0777)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
		log.Fatal(err)
	}

	if !info.IsDir() {
		log.Fatal("Path is not directory. Remove or replace file - " + path)
	}
}

func ensureFile(path string) {
	info, err := os.Stat(path)
	if err != nil {
		if err.Error() == cerrors.NewStatNoSuchFileErr(path).Error() {
			file, err := os.OpenFile(path, os.O_CREATE, 0777)
			if err != nil {
				log.Fatal(err)
			}
			file.Close()
			return
		}
		log.Fatal(err)
	}

	if info.IsDir() {
		log.Fatal("Path is directory. Remove or replace file - " + path)
	}
}

func getSizeAsString(size int64) string {
	mbs := size / 1024 / 1024
	if mbs < 1 {
		return "less than 1 MB"
	}

	return fmt.Sprintf("%d MB", mbs)
}
