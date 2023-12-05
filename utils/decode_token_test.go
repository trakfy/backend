package utils

import "testing"

func TestDecodeToken(t *testing.T) {
	t.Setenv("JWT_SECRET", "secret")
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImVtYWlsQG1haWwuY29tIiwiaWF0IjoxNjg0NzA2NTEwLCJleHAiOjEwMDAxNjg0NzA2NTA5fQ.nPgPksgu7OsTOXp2CBTJfPpYcewUoHF6blcM-n3wI3c"
	expectedEmail := "email@mail.com"

	claims, err := DecodeToken(token)
	if err != nil {
		t.Errorf("Expected nil, but got %v", err)
	}

	if claims.Email != expectedEmail {
		t.Errorf("Expected email %s, but got %s", expectedEmail, claims.Email)
	}
}
