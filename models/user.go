package models

type User struct {
	Sub  string `json:"sub" binding:"required"`
	Name string `json:"name" binding:"required"`
}
