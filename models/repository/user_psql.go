package repository

import (
	"lensent/db"
	"lensent/models"

	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

type User models.User

// GetAll is getting all Users
func (_ UserRepository) GetAll() ([]User, error) {
	db := db.ConnectDB()
	var u []User
	if err := db.Table("users").Select("sub,name").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// CreateModel is creating User model
func (_ UserRepository) CreateModel(c *gin.Context) (User, error) {
	db := db.ConnectDB()
	var u User
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	if err := db.Create(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// GetBySub is getting a User by sub
func (_ UserRepository) GetBySub(sub string) (models.User, error) {
	db := db.ConnectDB()
	var u models.User
	if err := db.Table("users").Where("sub = ?", sub).Scan(&u).Error; err != nil {
		return u, err
	}
	return u, nil
}

// UpdateBySub is update a User
func (_ UserRepository) UpdateBySub(sub string, c *gin.Context) (models.User, error) {
	db := db.ConnectDB()
	var u models.User
	if err := db.Table("users").Where("sub = ?", sub).Scan(&u).Error; err != nil {
		return u, err
	}
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	db.Save(&u)

	return u, nil
}

// DeleteBySub is delete a User by Sub
func (_ UserRepository) DeleteBySub(sub string) error {
	db := db.ConnectDB()
	var u User

	if err := db.Table("users").Where("sub = ?", sub).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
