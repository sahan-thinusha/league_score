package model

import "github.com/jinzhu/gorm"

type Match struct {
	Model
	Path           string      `json:"path"`
	Uid int64   `json:"uid"`
}


func (Match) TableName() string {
	return "match"
}
func (m *Match) PreloadAvailability(db *gorm.DB) *gorm.DB {
	return db.Preload("match")
}