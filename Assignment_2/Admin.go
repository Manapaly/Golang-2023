package main

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

type Service struct {
}

var Admin Service

func (Service) CreateUser(newUser User) error {
	Deserialize()
	_, ok := DB[newUser.Email]

	if ok {
		return errors.New("user already exists")
	}

	hash, err := getPasswordHash(newUser.Password)
	if err != nil {
		return nil
	}
	newAuthUser := AuthUser{newUser.Email, hash}
	DB[newUser.Email] = newAuthUser

	// serialiing all created users to json-database
	Serialize()

	return nil
}

func (Service) VerifyUser(user User) bool {
	Deserialize()
	authUser, ok := DB[user.Email]

	if !ok {
		return false
	}

	err := bcrypt.CompareHashAndPassword([]byte(authUser.PasswordHash), []byte(user.Password))
	return err == nil
}

func (Service) ViewUsers() {
	// deserialzing all users from json file
	Deserialize()

	for key, val := range DB {
		fmt.Println(key, val.PasswordHash)
	}
}
