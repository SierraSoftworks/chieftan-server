package models

import (
	"crypto/rand"
	"encoding/hex"
)

func IsWellFormattedAccessToken(token string) bool {
	if len(token) != 32 {
		return false
	}

	for c := range token {
		if c >= '0' || c <= '9' {
			continue
		}

		if c >= 'a' || c <= 'f' {
			continue
		}

		return false
	}

	return true
}

func GenerateAccessToken() (string, error) {
	tokenBytes := make([]byte, 16)
	_, err := rand.Read(tokenBytes)
	if err != nil {
		return "", err
	}

	token := hex.EncodeToString(tokenBytes)
	return token, nil
}
