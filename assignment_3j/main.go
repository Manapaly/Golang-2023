package main

import (
	"github.com/ZhanserikKalmukhambet/assignment_3/controllers"
	"github.com/ZhanserikKalmukhambet/assignment_3/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
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
