package repositories

import (
	"deathtiny_encounters/models"
	"gorm.io/gorm"
)

var db *gorm.DB

func InitDB(database *gorm.DB) {
	db = database
	db.AutoMigrate(&models.User{})
}

func CreateUser(user *models.User) error {
	return db.Create(user).Error
}

func GetUserByEmail(email string) (models.User, error) {
	var user models.User
	err := db.Where("email = ?", email).First(&user).Error
	return user, err
}
