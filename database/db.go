package database

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/shivpalyadavv/week2-GL1-CipherSchools/models"
	"gorm.io/gorm"
)

var DB *gorm.DB

func GetDB() *gorm.DB {
	return DB
}

func Setup() {
	host := "localhost"
	port := "5432"
	dbName := "book"
	username := "postgres"
	password := "1111"
	args := "host=" + host + "port=" + port + "user=" + username + "dbname=" + dbName + "sslmode=disable password=" + password
	db, err := gorm.Open("postgres", args)
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(models.Book{})
	DB = db
}
