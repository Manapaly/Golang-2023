package main

import (
	"github.com/ZhanserikKalmukhambet/assignment_3/initializers"
	"github.com/ZhanserikKalmukhambet/assignment_3/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
}

func main() {
	initializers.DB.AutoMigrate(&models.Book{})
}
