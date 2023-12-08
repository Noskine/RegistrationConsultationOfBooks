package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func PasswordHash(p []byte) string {
	hashedPassword, err := bcrypt.GenerateFromPassword(p, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}

	return string(hashedPassword)
}
