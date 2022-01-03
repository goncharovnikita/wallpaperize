package main

import (
	"fmt"
	"os"
)

func ensureDir(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, 0777)
		if err != nil {
			return err
		}

		return nil
	}

	if !info.IsDir() {
		return fmt.Errorf("path is not directory. Remove or replace file - %s", path)
	}

	return nil
}

func ensureFile(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		file, err := os.OpenFile(path, os.O_CREATE, 0777)
		if err != nil {
			return fmt.Errorf("error opening file %s: %w", path, err)
		}

		file.Close()

		return nil
	}

	if info.IsDir() {
		return fmt.Errorf("path is directory. Remove or replace file - %s", path)
	}

	return nil
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
