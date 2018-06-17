package main

import (
	"fmt"
	"log"
	"os"
)

func ensureDir(path string) {
	info, err := os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, 0777)
		if err != nil {
			log.Fatal(err)
		}
		return
	}

	if !info.IsDir() {
		log.Fatal("Path is not directory. Remove or replace file - " + path)
	}
}

func ensureFile(path string) {
	info, err := os.Stat(path)
	if err != nil {
		file, err := os.OpenFile(path, os.O_CREATE, 0777)
		if err != nil {
			log.Fatal(err)
		}
		file.Close()
		return
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

func getFileNames(path string, infos []os.FileInfo) []string {
	result := make([]string, 0, len(infos))
	for _, v := range infos {
		result = append(result, path+"/"+v.Name())
	}

	return result
}
