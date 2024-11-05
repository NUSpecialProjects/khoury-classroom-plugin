package utils

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerateToken(length uint32) (string, error) {
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	token := hex.EncodeToString(bytes)
	return token, nil
}
