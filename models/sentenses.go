package models

import (
	"time"
)

type Sentense struct {
	SentenseID string `json:"sentense_id" gorm:"primarykey" sql:"type:uuid"`
	Title      string `json:"title" binding:"required"`
	Sentense   string `json:"sentense" binding:"required"`
	UserID     string `json:"sub" binding:"required"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
}
