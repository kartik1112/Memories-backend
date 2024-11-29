package models

import (
	"errors"
	"time"

	"github.com/kartik1112/Memories-backend/initializers"
	"github.com/kartik1112/Memories-backend/utils"
)

type Event struct {
	ID        uint      `gorm:"primary_key;autoIncrement"`
	EventName string    `gorm:"not null" json:"event_name"`
	Date      time.Time `gorm:"not null" json:"date"`
	Code      string    `gorm:"size:6;not null;unique"`
	CreatedBy uint      `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`

	User User `gorm:"foreignKey:CreatedBy;constraint:onDelete:cascade" `
}

func (e *Event) CreateEvent() error {
	e.Code = utils.GenerateEventCode(e.CreatedBy, e.EventName)

	result := initializers.DB.Create(&e)

	if result.Error != nil {
		return result.Error
	}

	return nil

}

func (e *Event) GetEventFromCode() (*Event, error) {
	result := initializers.DB.Where("code = ?", e.Code).First(&e)
	if result.Error != nil {
		return nil, errors.New("Event Not Found with code" + e.Code)
	}
	return e, nil
}

func GetEventsCreatedByUserId(userId uint) ([]*Event, error) {
	var events []*Event
	result := initializers.DB.Preload("User").Where("created_by = ?", userId).Find(&events)
	if result.Error != nil {
		return nil, result.Error
	}
	return events, nil
}

func GetEventsJoinedByUserId(userId uint) ([]*Event, error) {
	var events []*Event
	result := initializers.DB.Joins("JOIN user_events ON user_events.event_id=events.id").Where("user_events.user_id = ?", userId).Find(&events)

	if result.Error != nil {
		return nil, result.Error
	}

	return events, nil
}
