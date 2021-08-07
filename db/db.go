package db

import (
	"lensent/models"

	"github.com/jinzhu/gorm"

	_ "github.com/lib/pq"
)

var (
	db  *gorm.DB
	err error
)

func Init() {
	db, err = gorm.Open("postgres", "host=PostgresTest port=5432 user=postgres password=postgres dbname=testdb sslmode=disable")
	if err != nil {
		panic("Database couldn't connect.")
	}
	autoMigration()
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
	db.AutoMigrate(&models.Sentense{}).AddForeignKey("user_id", "users(sub)", "RESTRICT", "RESTRICT")
	db.AutoMigrate(&models.Word{}).AddForeignKey("user_id", "users(sub)", "RESTRICT", "RESTRICT").AddForeignKey("sentense_id", "sentenses(sentense_id)", "RESTRICT", "RESTRICT")
}
