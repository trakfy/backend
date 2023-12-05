package utils

import (
	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	return uuid
}
