package models

import (
	"time"

	"github.com/kartik1112/Memories-backend/initializers"
)

type User struct {
	ID           uint   `gorm:"primary_key;autoIncrement"`
	Name         string `gorm:"not null" json:"name"`
	Email        string `gorm:"unique;not null" json:"email"`
	Password     string `gorm:"not null" json:"password"`
	UserImageUrl string `json:"userimageurl"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) CreateUser() error {
	result := initializers.DB.Create(&u)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (u *User) GetUserByEmail() error {
	result := initializers.DB.Where("email = ?", u.Email).First(&u)
	if result.Error!=nil {
		return result.Error
	}
	return nil
}
