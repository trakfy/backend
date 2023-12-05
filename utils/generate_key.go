package utils

import (
	"crypto/sha256"
	"encoding/hex"
)

func GenerateSHA256() string {
	value := GenerateUUID().String()
	hash := sha256.Sum256([]byte(value))
	return hex.EncodeToString(hash[:])
}
