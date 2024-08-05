package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	UserID      uint   `gorm:"unique;not null"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Age         int
	Gender      string
	Interests   string
	Description string
}
