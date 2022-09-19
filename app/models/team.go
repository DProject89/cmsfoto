package models

import (
	"time"

	"gorm.io/gorm"
)

type Team struct {
	ID                 string `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Name               string `gorm:"size:255;not null"`
	PicTeam            string `gorm:"size:255"`
	PicTeamPhoneNumber string `gorm:"size:255"`
	PicTeamEmail       string `gorm:"size:255"`
	InstagramTeam      string `gorm:"size:255"`
	Location           string `gorm:"size:255"`
	City               string `gorm:"size:255"`
	LogoTeam           string `gorm:"size:255"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

func (p *Team) GetTeams(db *gorm.DB) (*[]Team, error) {
	var err error
	var teams []Team

	err = db.Debug().Model(&Team{}).Limit(20).Find(&teams).Error
	if err != nil {
		return nil, err
	}

	return &teams, nil
}
