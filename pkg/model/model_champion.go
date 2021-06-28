package model

import "github.com/jinzhu/gorm"

type Champion struct {
	Model
	Name           string      `json:"name"`
	Index           string      `json:"index" gorm:"column:ch_id"`
	Group    string      `json:"group"`
	Avatar string  `json:"avatar"`
}

func (Champion) TableName() string {
	return "champion"
}
func (m *Champion) PreloadPatient(db *gorm.DB) *gorm.DB {
	return db
}