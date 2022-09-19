package models

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	ID          int    `gorm:"not null;uniqueIndex;primary_key"`
	Role        string `gorm:"size:255"`
	Description string `gorm:"size:255"`
	StatusRole  int    `gorm:"default:0"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func (p *Role) GetRoles(db *gorm.DB) (*[]Role, error) {
	var err error
	var roles []Role

	err = db.Debug().Model(&Role{}).Limit(20).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return &roles, nil
}
