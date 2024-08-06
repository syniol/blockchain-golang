package cryptography

import (
	"crypto/ed25519"
	"errors"
)

func NewSignature(encryptedData []byte, privateKey ed25519.PrivateKey) ([]byte, error) {
	hashedValue, err := NewHash(encryptedData)
	if err != nil {
		return nil, err
	}

	return ed25519.Sign(privateKey, hashedValue), nil
}

func VerifySignature(encryptedData, sig []byte, publicKey ed25519.PublicKey) (bool, error) {
	hashedValue, err := NewHash(encryptedData)
	if err != nil {
		return false, errors.New("error creating  hash from payload")
	}

	return ed25519.Verify(publicKey, hashedValue, sig), nil
}
