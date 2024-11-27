package models

import (
	"time"
)

type User struct {
	UserId         uint   `gorm:"primary_key;autoIncrement"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	UserImageUrl string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
