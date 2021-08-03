package db

import (
	"lensent/models"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var (
	db *gorm.DB
)

func Init() {
	db, err := gorm.Open("postgres", "host=PostgresTest port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		panic("Database couldn't connect.")
	}
	autoMigration()
	defer db.Close()
}

func ConnectDB() *gorm.DB {
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
}
