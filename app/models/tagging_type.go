package models

import (
	"time"

	"gorm.io/gorm"
)

type TaggingType struct {
	ID        int `gorm:"size:36;not null;uniqueIndex;primary_key"`
	Code      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func (p *TaggingType) GetTaggingTypes(db *gorm.DB) (*[]TaggingType, error) {
	var err error
	var taggingTypes []TaggingType

	err = db.Debug().Model(&TaggingType{}).Limit(20).Find(&taggingTypes).Error
	if err != nil {
		return nil, err
	}

	return &taggingTypes, nil
}
