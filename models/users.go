package models

import (
	"time"
)

type User struct {
	Sub        string `json:"sub" binding:"required" gorm:"primarykey"`
	Name       string `json:"name" binding:"required"`
	Created_at time.Time
	Updated_at time.Time
	Deleted_at *time.Time
}
