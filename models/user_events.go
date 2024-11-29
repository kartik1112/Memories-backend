package models

import (
	"time"
)

type UserEvent struct {
	ID       uint      `gorm:"primary_key;auto_increment"`
	UserID   uint      `gorm:"not null"`
	EventID  uint      `gorm:"not null"`
	JoinedAt time.Time `gorm:"autoCreateTime"`
	User     User      `gorm:"foreignKey:UserID;constraint:onDelete:cascade"`
	Event    Event     `gorm:"foreignKey:EventID"`
}

type EventCode struct {
	Code string `json:"event_code" binding:"required"`
}
