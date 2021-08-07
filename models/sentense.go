package models

import (
	"time"
)

type Sentense struct {
	SentenseID string `json:"sentense_id" gorm:"unique;primarykey" sql:"type:uuid"`
	Title      string `json:"title" binding:"required" gorm:"not null"`
	Sentense   string `json:"sentense" binding:"required" gorm:"not null"`
	UserID     string `json:"sub" binding:"required" gorm:"not null"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
}
