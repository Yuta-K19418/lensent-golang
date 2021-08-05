package repository

import (
	"lensent/db"
	"lensent/models"

	"github.com/gin-gonic/gin"
)

type UserRepository struct{}

type User models.User

// GetAll is get all User
func (_ UserRepository) GetAll() ([]User, error) {
	db := db.ConnectDB()
	var u []User
	if err := db.Table("users").Select("id,name").Scan(&u).Error; err != nil {
		return nil, err
	}
	return u, nil
}

// CreateModel is create User model
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

// GetByid is getting a User by id
func (_ UserRepository) GetByID(id int) (models.User, error) {
	db := db.ConnectDB()
	var me models.User
	if err := db.Where("id = ?").First(&me).Error; err != nil {
		return me, err
	}
	return me, nil
}

// UpdateByid is update a User
func (_ UserRepository) UpdateByID(id int, c *gin.Context) (models.User, error) {
	db := db.ConnectDB()
	var u models.User
	if err := db.Where("id = ?", id).First(&u).Error; err != nil {
		return u, err
	}
	if err := c.BindJSON(&u); err != nil {
		return u, err
	}
	u.ID = uint(id)
	db.Save(&u)

	return u, nil
}

// DeleteByID is delete a User by ID
func (_ UserRepository) DeleteByID(id int) error {
	db := db.ConnectDB()
	var u User

	if err := db.Where("id = ?", id).Delete(&u).Error; err != nil {
		return err
	}

	return nil
}
