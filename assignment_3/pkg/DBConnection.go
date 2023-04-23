package pkg

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() (gorm.DB, error) {
	//err1 := godotenv.Load()
	//if err1 != nil {
	//	panic(err1)
	//}
	//db, err := gorm.Open(postgres.Open(os.Getenv("host=localhost user=postgres password=321qaz dbname=BookStore port=5432")), &gorm.Config{})
	db, err := gorm.Open(postgres.Open("host=db user=postgres password=321qaz dbname=BookStore port=5432"), &gorm.Config{})

	if err != nil {
		panic(err)
	} else {
		return *db, nil
	}
}
