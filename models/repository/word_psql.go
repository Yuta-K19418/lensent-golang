package repository

import (
	"lensent/db"
	"lensent/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type WordRepository struct{}

type Word models.Word

// GetAll is getting all Words
func (_ WordRepository) GetAll() ([]Word, error) {
	db := db.ConnectDB()
	var w []Word
	if err := db.Table("words").Select("word_id, en, ja, user_id, sentense_id").Scan(&w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

// CreateModel is creating Word Model
func (_ WordRepository) CreateModel(c *gin.Context) (Word, error) {
	db := db.ConnectDB()
	var w Word
	w.WordID = uuid.NewString()
	if err := c.BindJSON(&w); err != nil {
		return w, err
	}
	if err := db.Create(&w).Error; err != nil {
		return w, err
	}
	return w, nil
}

// GetAllByUser is getting all Words by Sentense
func (_ WordRepository) GetAllBySentense(sentense_id string) ([]Word, error) {
	db := db.ConnectDB()
	var w []Word
	if err := db.Table("words").Where("sentense_id = ?", sentense_id).Scan(&w).Error; err != nil {
		return nil, err
	}
	return w, nil
}

// GetByWordId is getting a Word by word_id
func (_ WordRepository) GetByWordId(word_id string) (models.Word, error) {
	db := db.ConnectDB()
	var w models.Word
	if err := db.Table("words").Where("word_id = ?", word_id).Scan(&w).Error; err != nil {
		return w, err
	}
	return w, nil
}

// UpdateByWordId is update a Word
func (_ WordRepository) UpdateByWordId(word_id string, c *gin.Context) (models.Word, error) {
	db := db.ConnectDB()
	var w models.Word
	if err := db.Where("word_id = ?", word_id).First(&w).Error; err != nil {
		return w, err
	}
	if err := c.BindJSON(&w); err != nil {
		return w, err
	}
	db.Save(&w)

	return w, nil
}

// DeleteByWordId is delete a Word by word_id
func (_ WordRepository) DeleteByWordId(word_id string) error {
	db := db.ConnectDB()
	var w Word

	if err := db.Table("words").Where("word_id = ?", word_id).Delete(&w).Error; err != nil {
		return err
	}

	return nil
}
