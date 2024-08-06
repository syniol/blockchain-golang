package cryptography

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
)

func NewHMAC(value, secret []byte) ([]byte, error) {
	if secret == nil {
		var err error

		secret, err = NewHMACSaltKey()
		if err != nil {
			return nil, err
		}
	}

	hmacHash := hmac.New(sha512.New, secret)
	_, err := hmacHash.Write(value)
	if err != nil {
		return nil, err
	}

	output := make([]byte, hex.EncodedLen(hmacHash.Size()))
	hex.Encode(output, hmacHash.Sum(nil))

	return output, nil
}

func NewHMACSaltKey() ([]byte, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(randomBytes)))

	base64.StdEncoding.Encode(dst, randomBytes)

	return dst, nil
}
