package model

import "github.com/jinzhu/gorm"

type Gamer struct {
	Model
	Name           string      `json:"name"`
	Gender         string      `json:"gender"`
	DateOfBirth    string      `json:"dob"`
	ProfilePic string  `json:"profilepic"`
	RiotId  string      `json:"riotId"`
	Region  string      `json:"region"`
	User     *User `gorm:"foreignkey:userID" json:"user"`
	UserID int64
}

func (Gamer) TableName() string {
	return "gamer"
}
func (m *Gamer) PreloadPatient(db *gorm.DB) *gorm.DB {
	return db
}
