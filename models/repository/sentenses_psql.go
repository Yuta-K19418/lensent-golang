package repository

import (
	"lensent/db"
	"lensent/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type SentensesRepository struct{}

type Sentense models.Sentense

// GetAll is getting all sentenses
func (_ SentensesRepository) GetAll() ([]Sentense, error) {
	db := db.ConnectDB()
	var s []Sentense
	if err := db.Table("sentenses").Select("sentense_id, title, sentense, users").Scan(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// CreateModel is creating Sentense Model
func (_ SentensesRepository) CreateModel(c *gin.Context) (Sentense, error) {
	db := db.ConnectDB()
	var s Sentense
	s.SentenseID = uuid.NewString()
	if err := c.BindJSON(&s); err != nil {
		return s, err
	}
	if err := db.Create(&s).Error; err != nil {
		return s, err
	}
	return s, nil
}

// GetAllByUser is getting all sentenses by User
func (_ SentensesRepository) GetAllByUser(sub string) ([]Sentense, error) {
	db := db.ConnectDB()
	var s []Sentense
	if err := db.Table("sentenses").Where("sub = ?").Select("sentense_id, title, sentense, users").Scan(&s).Error; err != nil {
		return nil, err
	}
	return s, nil
}

// GetBySentenseId is getting a Sentense by sentense_id
func (_ SentensesRepository) GetBySentenseId(sentense_id string) (models.Sentense, error) {
	db := db.ConnectDB()
	var s models.Sentense
	if err := db.Where("sentense_id = ?").First(&s).Error; err != nil {
		return s, err
	}
	return s, nil
}

// UpdateBySentenseId is update a Sentense
func (_ SentensesRepository) UpdateBySentenseId(sentense_id string, c *gin.Context) (models.Sentense, error) {
	db := db.ConnectDB()
	var s models.Sentense
	if err := db.Where("sentense_id = ?", sentense_id).First(&s).Error; err != nil {
		return s, err
	}
	if err := c.BindJSON(&s); err != nil {
		return s, err
	}
	s.SentenseID = sentense_id
	db.Save(&s)

	return s, nil
}

// DeleteBySentenseId is delete a Sentense by sentense_id
func (_ SentensesRepository) DeleteBySentenseId(sentense_id string) error {
	db := db.ConnectDB()
	var s Sentense

	if err := db.Where("sentense_id = ?", sentense_id).Delete(&s).Error; err != nil {
		return err
	}

	return nil
}
