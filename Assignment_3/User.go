package main

import "golang.org/x/crypto/bcrypt"

type User struct {
	FirsName string `json:"firs_name"`
	LastName string `json:"last_name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUser struct {
	Email        string
	PasswordHash string
}

func getPasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 0)
	return string(hash), err // hash is also considered as a string
}
