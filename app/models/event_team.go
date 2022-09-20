package models

import (
	"time"

	"gorm.io/gorm"
)

type EventTeam struct {
	ID           int    `gorm:"not null;uniqueIndex;primary_key"`
	Name         string `gorm:"size:255;not null"`
	Code         string `gorm:"size:255"`
	EventID      int
	EventGroupID int
	TeamID       int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

func (p *EventTeam) GetEventTeams(db *gorm.DB) (*[]EventTeam, error) {
	var err error
	var EventTeams []EventTeam

	err = db.Debug().Model(&EventTeam{}).Limit(20).Find(&EventTeams).Error
	if err != nil {
		return nil, err
	}

	return &EventTeams, nil
}

func (p *EventTeam) GetEventTeamsByGroup(db *gorm.DB, eventGroupId string) (*[]EventTeam, error) {
	var err error
	var EventTeams []EventTeam

	err = db.Debug().Model(&EventTeam{}).Where("event_group_id = ?", eventGroupId).Find(&EventTeams).Error
	if err != nil {
		return nil, err
	}

	return &EventTeams, nil
}
