package model


import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Model
	Name string  `json:"name"`
	Email  string  `json:"email"`
	Password  string  `json:"password"  gorm:"size:2000"`
	Token     string  `json:"token"`
	Role     string  `json:"role"`
}

func (m *User) PreloadUser(db *gorm.DB) *gorm.DB {
	return db
}

