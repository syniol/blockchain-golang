package cryptography

import (
	"testing"
)

// TestNewHash expected values for hash been created by https://emn178.github.io/online-tools/sha512.html
func TestNewHash(t *testing.T) {
	// Given
	sut := NewHash

	t.Run("when a string provided as a parameter", func(t *testing.T) {
		actual, _ := sut([]byte("Syniol Limited"))
		expected := "31d26396c540a83a7bd04e925255c774096a0edfd5312556eaae861b087b12c8a446a5e42bf240a9a994b7b2d5aebf86b93c45ec219eb23a8ad8004cf353ed8b"

		// Then
		if string(actual) != expected {
			t.Fatalf("it was expecting %s but got %s", expected, actual)
		}
	})

	t.Run("when an empty string provided as a parameter", func(t *testing.T) {
		actual, _ := sut([]byte(""))
		expected := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"

		// Then
		if string(actual) != expected {
			t.Fatalf("it was expecting %s but got %s", expected, actual)
		}
	})

	t.Run("when a nil provided as a parameter", func(t *testing.T) {
		actual, _ := sut(nil)
		expected := "cf83e1357eefb8bdf1542850d66d8007d620e4050b5715dc83f4a921d36ce9ce47d0d13c5d85f2b0ff8318d2877eec2f63b931bd47417a81a538327af927da3e"

		// Then
		if string(actual) != expected {
			t.Fatalf("it was expecting %s but got %s", expected, actual)
		}
	})
}
