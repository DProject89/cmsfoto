package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID                 int    `gorm:"not null;uniqueIndex;primary_key"`
	Name               string `gorm:"size:255;not null"`
	Code               string `gorm:"size:255"`
	PicTeam            string `gorm:"size:255"`
	PicTeamPhoneNumber string `gorm:"size:255"`
	PicTeamEmail       string `gorm:"size:255"`
	InstagramTeam      string `gorm:"size:255"`
	Location           string `gorm:"size:255"`
	LogoTeam           string `gorm:"size:255"`
	City               string `gorm:"size:255"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

func (p *Team) GetTeams(db *gorm.DB) (*[]Team, error) {
	var err error
	var Teams []Team

	err = db.Debug().Model(&Team{}).Limit(20).Find(&Teams).Error
	if err != nil {
		return nil, err
	}

	return &Teams, nil
}
