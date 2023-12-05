package utils

import "testing"

func TestGenerateUUID(t *testing.T) {
	expectedLength := 36

	result := GenerateUUID()

	string_uuid := result.String()

	if len(string_uuid) != expectedLength {
		t.Errorf("Expected length %d, but got %d", expectedLength, len(string_uuid))
	}
}
