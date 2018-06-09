package hash

import (
	"crypto/sha256"
	"fmt"
)

// Hash256 hashes given bytes to string
func Hash256(r []byte) string {
	hsh := sha256.New()

	hsh.Write(r)

	return fmt.Sprintf("%x", hsh.Sum(nil))
}
