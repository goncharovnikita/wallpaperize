package hash

import (
	"crypto/sha256"
	"fmt"
)

// Hash256 hashes given bytes to string
func Hash256(r []byte) (string, error) {
	hsh := sha256.New()

	if _, err := hsh.Write(r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hsh.Sum(nil)), nil
}
