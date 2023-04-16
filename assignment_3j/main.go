package main

import (
	"github.com/Manapaly/Golang-2023/assignment_3j/controllers"
	"github.com/Manapaly/Golang-2023/assignment_3j/initializers"
	"github.com/Manapaly/Golang-2023/assignment_3j/models"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.DB.AutoMigrate(&models.Book{})
}

func main() {
	r := gin.Default()

	r.GET("/books", controllers.GetBooks)
	r.GET("/books/id/:id", controllers.GetBookByID)
	r.GET("/books/title/:title", controllers.GetBookByTitle)
	r.GET("/booksByPriceInAsc", controllers.GetBookSortedByCostAsc)
	r.GET("/booksByPriceInDesc", controllers.GetBookSortedByCostDesc)

	r.POST("/books", controllers.CreateBook)

	r.PATCH("/books/:id", controllers.UpdateBook)

	r.DELETE("/books/:id", controllers.DeleteBook)

	r.Run()
}
