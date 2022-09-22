package models

import (
	"time"

	"gorm.io/gorm"
)

type Tagging struct {
	ID            int `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Code          string
	Name          string
	TaggingTypeId int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (p *Tagging) GetTaggings(db *gorm.DB) (*[]Tagging, error) {
	var err error
	var taggings []Tagging

	err = db.Debug().Model(&Tagging{}).Limit(20).Find(&taggings).Error
	if err != nil {
		return nil, err
	}

	return &taggings, nil
}
