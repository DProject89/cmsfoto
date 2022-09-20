package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	ID             int `gorm:"not null;uniqueIndex;primary_key"`
	Code           string
	Name           string
	Location       string
	City           string
	Description    string
	DateStart      time.Time
	DateFinish     time.Time
	TimeStart      string
	TimeFinish     string
	EventGroupId   int
	PicEvent       string
	TeamAId        int
	TeamBId        int
	PaymentMethod  string
	Price          int
	PhotographerId int
	CoverImage     string
	LinkYotube     string
	LinkMedia      string
	CreatedBy      int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      gorm.DeletedAt
}

func (p *Event) GetEvents(db *gorm.DB) (*[]Event, error) {
	var err error
	var Events []Event

	err = db.Debug().Model(&Event{}).Limit(20).Find(&Events).Error
	if err != nil {
		return nil, err
	}

	return &Events, nil
}
