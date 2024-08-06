package cryptography

import (
	"crypto/ed25519"
	"crypto/rand"
	"testing"
)

func TestNewSignature(t *testing.T) {
	// Given
	sut := NewSignature
	message := []byte("Syniol Limited")

	pub, prv, _ := ed25519.GenerateKey(rand.Reader)

	t.Run("When is executed with a valid value", func(t *testing.T) {

		//Then
		actual, err := sut(message, prv)

		if err != nil {
			t.Errorf("it was not expecting an error but got %s", err.Error())
		}

		t.Run("It should create and verify the signature", func(t *testing.T) {
			if resp, _ := VerifySignature(message, actual, pub); resp == false {
				t.Error("failed to verify digital signature")
			}
		})
	})
}
