package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `gorm:"not null"`
	Email        string `gorm:"unique;not null"`
	PasswordHash string `gorm:"not null"`
	UserImageUrl string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
