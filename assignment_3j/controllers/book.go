package controllers

import (
	"github.com/ZhanserikKalmukhambet/assignment_3/initializers"
	"github.com/ZhanserikKalmukhambet/assignment_3/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func GetBooks(c *gin.Context) {
	var books []models.Book

	initializers.DB.Find(&books)

	if len(books) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book store is empty. No book added!"})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": books})
	}
}

func GetBookByID(c *gin.Context) {
	var book models.Book

	id := c.Param("id")
	initializers.DB.First(&book, id)

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "No such ID"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func GetBookSortedByCostAsc(c *gin.Context) {
	var sorted []models.Book

	initializers.DB.Find(&sorted)

	if len(sorted) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book store is empty. No book added!"})
	} else {
		initializers.DB.Order("cost").Find(&sorted)
		c.JSON(http.StatusOK, gin.H{"sorted by price in ascending order": sorted})
	}
}

func GetBookSortedByCostDesc(c *gin.Context) {
	var sorted []models.Book

	initializers.DB.Find(&sorted)

	if len(sorted) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Book store is empty. No book added!"})
	} else {
		initializers.DB.Order("cost desc").Find(&sorted)
		c.JSON(http.StatusOK, gin.H{"sorted by price in ascending order": sorted})
	}
}

func GetBookByTitle(c *gin.Context) { // find by snakecase letter
	var book models.Book
	title := c.Param("title")
	res := strings.ToLower(title)

	initializers.DB.Where("LOWER(REPLACE(title, ' ', '')) = ?", res).Find(&book)

	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"found": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"found": book})
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	initializers.DB.First(&book, id)

	var input models.UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	initializers.DB.Model(&book).Updates(input)
	c.JSON(http.StatusOK, gin.H{"book is updated": book})
}

func DeleteBook(c *gin.Context) {
	var book models.Book

	id := c.Param("id")

	initializers.DB.First(&book, id)
	if book.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no such book to delete"})
		return
	}

	initializers.DB.Delete(&book)
	c.JSON(http.StatusBadRequest, gin.H{"deleted": true})
}

func CreateBook(c *gin.Context) {
	var input models.CreateBookInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	book := models.Book{
		Title:       input.Title,
		Author:      input.Author,
		Description: input.Description,
		Cost:        input.Cost,
	}

	initializers.DB.Create(&book)
	c.JSON(http.StatusOK, gin.H{"new book added": book})
}
