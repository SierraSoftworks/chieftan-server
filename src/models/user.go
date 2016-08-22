package models

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

type UserSummary struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type User struct {
	ID          string   `json:"id" bson:"_id"`
	Name        string   `json:"name"`
	Email       string   `json:"email"`
	Permissions []string `json:"permissions,omitempty"`
	Tokens      []string `json:"tokens,omitempty"`
}

func (u *User) UpdateID() {
	u.ID = DeriveID(u.Email)
}

func (u *User) Summary() UserSummary {
	return UserSummary{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}

func (u *User) NewToken() (string, error) {
	token, err := GenerateAccessToken()
	if err != nil {
		return token, err
	}
	u.Tokens = append(u.Tokens, token)

	return token, nil
}

func DeriveID(email string) string {
	normalizedEmail := strings.ToLower(strings.TrimSpace(email))
	emailData := []byte(normalizedEmail)

	hash := md5.Sum(emailData)
	return hex.EncodeToString(hash[:])
}

func IsValidUserID(id string) bool {
	if len(id) != 32 {
		return false
	}

	for c := range id {
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
