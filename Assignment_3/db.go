package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

var DB = map[string]AuthUser{}

// key = string->email
// val = authUser->{email, passwordHash}

func Serialize() {
	byteValue, _ := json.MarshalIndent(DB, "", " ")
	_ = ioutil.WriteFile("users.json", byteValue, 0644)
}

func Deserialize() {
	jsonFile, err := os.Open("users.json")

	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	json.Unmarshal([]byte(byteValue), &DB)
}
