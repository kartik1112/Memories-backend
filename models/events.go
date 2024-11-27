package models

import (
	"time"
)

type Events struct {
	EventID   uint      `gorm:"primary_key;autoIncrement"`
	EventName string    `gorm:"not null" json:"event_name"`
	Date      time.Time `gorm:"not null" json:"date"`
	Code      string    `gorm:"size:6;not null;unique"`
	CreatedBy uint    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:CreatedBy;constraint:onDelete:cascade"`
}
