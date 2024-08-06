package cryptography

import (
	"crypto/sha512"
	"encoding/hex"
)

// NewHash will generate SHA-512 hash from value given as a parameter
func NewHash(value []byte) ([]byte, error) {
	hash := sha512.New()
	_, err := hash.Write(value)
	if err != nil {
		return nil, err
	}

	dist := make([]byte, hex.EncodedLen(hash.Size()))
	hex.Encode(dist, hash.Sum(nil))

	return dist, nil
}
