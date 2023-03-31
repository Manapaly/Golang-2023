package main

import (
	"github.com/Manapaly/Golang-2023/assignment_3j/initializers"
	"github.com/Manapaly/Golang-2023/assignment_3j/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}
