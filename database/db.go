package database

import (
	"log"

	//_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/shivpalyadavv/week2-GL1-CipherSchools/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Setup() {
	dsn := "host=localhost user=postgress password=1111 dbname=book port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	//db.AutoMigrate(models.Users{})
	DB = db
}
