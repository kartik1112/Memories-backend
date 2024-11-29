package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func GenerateHashedPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatal("Can not Generate Hash : ", err)
		return "", errors.New("Can not Generate Hash")
	}

	return string(hash), nil
}

func CheckPasswordAndHash(dbPassword string, userPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(dbPassword), []byte(userPassword))
	return err == nil
}
