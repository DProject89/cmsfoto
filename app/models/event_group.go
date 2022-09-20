package models

import (
	"time"

	"gorm.io/gorm"
)

type EventGroup struct {
	ID            int    `gorm:"not null;uniqueIndex;primary_key"`
	Code          string `gorm:"size:255"`
	Name          string `gorm:"size:255;not null"`
	Location      string `gorm:"size:255"`
	DateStart     time.Time
	DateFinish    time.Time
	PicOrganizer  string `gorm:"size:255"`
	OrganizerName string `gorm:"size:255;not null;uniqueIndex"`
	City          string `gorm:"size:255"`
	EventType     string `gorm:"size:255"`
	Price         string `gorm:"size:255"`
	EventLogo     string `gorm:"size:255"`
	HasEvent      int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (p *EventGroup) GetEventGroups(db *gorm.DB) (*[]EventGroup, error) {
	var err error
	var EventGroups []EventGroup

	err = db.Debug().Model(&EventGroup{}).Limit(20).Find(&EventGroups).Error
	if err != nil {
		return nil, err
	}

	return &EventGroups, nil
}
