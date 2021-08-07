package models

import (
	"time"
)

type User struct {
	Sub        string `json:"sub" binding:"required" gorm:"unique;primarykey;not null"`
	Name       string `json:"name" binding:"required" gorm:"not null"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
}
