package crawler

import (
	"os"
	"path/filepath"
)

func getDirSize(path string) (int64, error) {
	result := int64(0)

	err := filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if !info.IsDir() {
			result += info.Size()
		}

		return nil
	})

	return result, err
}

func getFilenamesFromDir(path string) ([]string, error) {
	fnames := make([]string, 0)
	err := filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}

		fnames = append(fnames, info.Name())
		return nil
	})

	return fnames, err
}
