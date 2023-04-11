package initializers

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func ConnectDatabase() {
	DB, err = gorm.Open(postgres.Open("host=db user=postgres password=321qaz dbname=BookStore port=5432"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to DATABASE")
	}
}
