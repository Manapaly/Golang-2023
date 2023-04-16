package main

import (
	"BookStore/pkg"
	"github.com/gin-gonic/gin"
)

func init() {
	db, _ := pkg.ConnectDB()
	db.AutoMigrate(&pkg.Book{})
}

func main() {
	r := gin.Default()
	r.GET("/books", pkg.GetBooks)
	r.GET("/books/desc", pkg.Desc)
	r.GET("/books/asc", pkg.Asc)
	r.GET("/books/:id", pkg.GetBook)
	r.POST("/books", pkg.CreateBook)
	r.DELETE("/books/:id", pkg.DeleteBook)
	r.PATCH("/books/:id", pkg.UpdateBook)
	r.Run()
}
