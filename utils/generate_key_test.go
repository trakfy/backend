package utils

import (
	"testing"
)

func TestGenerateSHA256(t *testing.T) {
	expectedLength := 64

	result := GenerateSHA256()

	if len(result) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(result))
	}
}
