package pkg

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	//ID          int    `gorm:"primaryKey"`
	Title       string `gorm:"column:title"`
	Author      string `gorm:"column:author"`
	Description string `gorm:"column:description"`
	Cost        int    `gorm:"column:cost"`
}
type UpdateBookInput struct {
	Title       string  `json:"title"`
	Author      string  `json:"author"`
	Description string  `json:"description"`
	Cost        float64 `json:"cost"`
}
