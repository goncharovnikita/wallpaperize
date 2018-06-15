package utils

import (
	"math/rand"
	"os"
	"path/filepath"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandStringBytes get random string
func RandStringBytes(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

// GetBytesFromGigabytes convert gigabytes to bytes
func GetBytesFromGigabytes(gb float64) int64 {
	result := gb * (1024 * 1024 * 1024)
	return int64(result)
}

// GetFilenamesFromDir get all filenames from directory
func GetFilenamesFromDir(path string) ([]string, error) {
	result := make([]string, 0)

	err := filepath.Walk(path, func(p string, info os.FileInfo, e error) error {
		if info.IsDir() {
			return nil
		}

		result = append(result, info.Name())
		return nil
	})

	return result, err
}
