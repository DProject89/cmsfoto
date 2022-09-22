package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID            int    `gorm:"size:36;not null;uniqueIndex;primary_key"`
	KTP           string `gorm:"size:255"`
	Name          string `gorm:"size:255;not null"`
	Gender        string `gorm:"size:255"`
	BirthPlace    string `gorm:"size:255"`
	BirthDate     time.Time
	Address       string `gorm:"size:255"`
	PhoneNumber   string `gorm:"size:255"`
	Email         string `gorm:"size:255;not null;uniqueIndex"`
	Password      string `gorm:"size:255"`
	RememberToken string `gorm:"size:255"`
	Instagram     string `gorm:"size:255"`
	Tiktok        string `gorm:"size:255"`
	Facebook      string `gorm:"size:255"`
	StatusId      int    `gorm:"default:0"`
	ImageProfile  string `gorm:"size:255"`
	MainTeamId    int
	TotalMedia    int `gorm:"default:0"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

func (p *User) GetUsers(db *gorm.DB) (*[]User, error) {
	var err error
	var users []User

	err = db.Debug().Model(&User{}).Limit(20).Find(&users).Error
	if err != nil {
		return nil, err
	}

	return &users, nil
}
