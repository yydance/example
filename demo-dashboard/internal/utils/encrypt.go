package utils

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
)

func EncodeHash(msg string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(msg+string(salt)), bcrypt.DefaultCost)
	return string(hash), err
}

func DecodeHash(hashMsg, msg string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashMsg), []byte(msg))
	return err == nil
}
