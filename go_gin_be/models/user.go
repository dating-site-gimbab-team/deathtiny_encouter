package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string  `gorm:"unique;not null"`
	Email    string  `gorm:"unique;not null"`
	Profile  Profile `gorm:"foreignkey:UserID"`
}
