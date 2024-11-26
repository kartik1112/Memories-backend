package models

import (
	"time"
)

type User struct {
	ID           uint   `gorm:"primary_key"`
	Name         string `json:"name" gorm:"not null"`
	Email        string `json:"email" gorm:"unique;not null"`
	PasswordHash string `json:"passwordhash" gorm:"not null"`
	Age          uint
	UserImageUrl string `json:"userimageurl"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
