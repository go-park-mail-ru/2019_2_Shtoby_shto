package utils

import (
	"fmt"
	"github.com/satori/go.uuid"
	"math/rand"
)

func GenerateUUID() (uuid.UUID, error) {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return uuid.Nil, err
	}
	return uuid.NewV4(), nil
}
