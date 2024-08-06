package cryptography

import "testing"

func TestNewHMAC(t *testing.T) {
	sut := NewHMAC

	t.Run("should create a same hash with a same key", func(t *testing.T) {
		// Given
		mockKey := []byte("5YQ_9hi4yD3QJ40yYS-T-A==")
		actual, _ := sut([]byte("MockValue"), mockKey)

		// Then
		expected := "ea843991af667a7af8e12cb069c3f1d712b41a180a515225fae852e62ce795675722584f5b9f44e847101accb69da815aa327f028b2d04314919f78d26205d1b"

		if string(actual) != expected {
			t.Errorf("it was exptecting %s but got %s", expected, string(actual))
		}
	})

	t.Run("should create a different hash when key is not provided", func(t *testing.T) {
		// Given
		actualFirstHash, _ := sut([]byte("MockValue"), nil)
		actualSecondHash, _ := sut([]byte("MockValue"), nil)

		// Then
		if string(actualFirstHash) == string(actualSecondHash) {
			t.Error(
				"it was not expecting a same hash",
				actualFirstHash,
				actualSecondHash,
			)
		}
	})
}
