package pkg

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func GetBooks(c *gin.Context) {
	db, _ := ConnectDB()
	var bks []Book
	if strings.Contains(c.Request.URL.String(), "/books?title=") {
		title := c.Request.URL.Query().Get("title")
		db.Model(&Book{}).Where("\"title\" = ?", title).Find(&bks)
		if len(bks) == 0 {
			c.JSON(404, gin.H{"message": "not found"})
		} else {
			c.JSON(http.StatusOK, bks)
		}
	} else {
		db.Find(&bks)
		c.JSON(http.StatusOK, bks)
	}
}

func GetBook(c *gin.Context) {
	db, _ := ConnectDB()
	id := c.Param("id")
	var bks []Book
	newId, _ := strconv.Atoi(id)
	db.Where("id = ?", newId).Find(&bks)
	if len(bks) == 0 {
		c.JSON(404, gin.H{"message": "not found!"})
	} else {
		c.JSON(200, bks)
	}
}

func CreateBook(c *gin.Context) {
	db, _ := ConnectDB()
	var b Book
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal([]byte(reqBody), &b)
	db.Create(&b)
	c.JSON(201, gin.H{"message": "created"})
}

func DeleteBook(c *gin.Context) {
	db, _ := ConnectDB()
	id := c.Param("id")
	var bks []Book
	newId, _ := strconv.Atoi(id)
	db.Where("id = ?", newId).Find(&bks)
	if len(bks) == 0 {
		c.JSON(404, gin.H{"message": "not found"})
	} else {
		db.Delete(&bks)
		c.JSON(200, gin.H{"message": "deleted"})
	}
}

func Desc(c *gin.Context) {
	db, _ := ConnectDB()
	var bks []Book
	db.Order("ID desc").Find(&bks)
	c.JSON(http.StatusOK, bks)
}

func Asc(c *gin.Context) {
	db, _ := ConnectDB()
	var bks []Book
	db.Order("ID asc").Find(&bks)
	c.JSON(http.StatusOK, bks)
}

func UpdateBook(c *gin.Context) {
	db, _ := ConnectDB()
	id := c.Param("id")
	reqBody, _ := ioutil.ReadAll(c.Request.Body)
	b := make(map[string]string)
	json.Unmarshal([]byte(reqBody), &b)
	IdNew, _ := strconv.Atoi(id)
	db.Model(&Book{}).Where("iD = ?", IdNew).Update("Title", b["Title"])
	db.Model(&Book{}).Where("iD = ?", IdNew).Update("Description", b["Description"])
	db.Model(&Book{}).Where("iD = ?", IdNew).Update("Author", b["Author"])
	db.Model(&Book{}).Where("iD = ?", IdNew).Update("Cost", b["Cost"])
	c.JSON(200, gin.H{"message": "updated"})
}

//func UpdateBook(c *gin.Context) {
//	db, _ := ConnectDB()
//	var b Book
//	id := c.Param("id")
//	if err := db.Where("id = ?", id).First(&b).Error; err != nil {
//		c.JSON(404, gin.H{"error": "book not found"})
//		return
//	}
//	reqBody, _ := ioutil.ReadAll(c.Request.Body)
//	json.Unmarshal([]byte(reqBody), &b)
//	db.Save(&b)
//	c.JSON(200, gin.H{"message": "updated"})
//}

//func UpdateBook(c *gin.Context) {
//	var book Book
//	db, _ := ConnectDB()
//	id := c.Param("id")
//	db.First(&book, id)
//
//	var input UpdateBookInput
//	if err := c.ShouldBindJSON(&input); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//
//	db.Model(&book).Updates(input)
//	c.JSON(http.StatusOK, gin.H{"book is updated": book})
//}
