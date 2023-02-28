package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Title  string `json:"title"`
	Author Author `json:"author"`
}

type Author struct {
	Sales     int  `json:"book_sales"`
	Age       int  `json:"age"`
	Developer bool `json:"is_developer"`
}

func main() {
	file, _ := ioutil.ReadFile("test.json")

	book := Book{}

	_ = json.Unmarshal([]byte(file), &book)
	fmt.Println(book.Title)
	//for i := 0; i < len(books); i++ {
	//	fmt.Println("Book title: ", books[i].Title)
	//}
}
