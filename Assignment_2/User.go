package main

import "golang.org/x/crypto/bcrypt"

type User struct {
	FirsName string
	LastName string
	Phone    string
	Email    string
	Password string
}

type AuthUser struct {
	Email        string
	PasswordHash string
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err // hash is also considered as a string
}
