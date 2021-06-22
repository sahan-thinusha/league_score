package model

import "github.com/jinzhu/gorm"

type ProMatches struct {
	Model
	Title string  `json:"title"`
	Description  string  `json:"description gorm:"size:2000""`
	VideoId  string  `json:"videoId"`
}

func (ProMatches) TableName() string {
return "pro_matches"
}
func (m *ProMatches) PreloadAvailability(db *gorm.DB) *gorm.DB {
	return db.Preload("ProMatches")
}