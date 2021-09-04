package db

import (
	"fmt"
	"lensent/models"
	"os"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	ConnectDB()
	autoMigration()
}

func ConnectDB() *gorm.DB {
	db, err = gorm.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_ENDPOINT"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME")))
	if err != nil {
		panic("Database couldn't connect.")
	}
	return db
}

func Close() {
	if err := db.Close(); err != nil {
		panic(err)
	}
}

func autoMigration() {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Sentense{}).AddForeignKey("user_id", "users(sub)", "CASCADE", "RESTRICT")
	db.AutoMigrate(&models.Word{}).AddForeignKey("user_id", "users(sub)", "CASCADE", "RESTRICT").AddForeignKey("sentense_id", "sentenses(sentense_id)", "CASCADE", "RESTRICT")
}
