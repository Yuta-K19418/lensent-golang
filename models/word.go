package models

import (
	"time"
)

type Word struct {
	WordID     string `json:"word_id" gorm:"unique;primarykey" sql:"type:uuid"`
	En         string `json:"en" bindging:"required" gorm:"not null"`
	Ja         string `json:"ja" bindging:"required" gorm:"not null"`
	UserID     string `json:"sub" binding:"required" gorm:"not null"`
	SentenseID string `json:"sentense_id" bindging:"required" gorm:"not null" sql:"type:uuid"`
	Created_at time.Time
	Updated_at time.Time
}
