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
