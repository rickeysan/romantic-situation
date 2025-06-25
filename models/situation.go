package models

import (
	"time"

	"gorm.io/gorm"
)

type Situation struct {
	ID          uint           `gorm:"primaryKey"`
	UserID      uint           `gorm:"not null"`
	User        User           `gorm:"foreignKey:UserID"`
	Title       string         `gorm:"not null"`
	Location    string         `gorm:"not null"`
	TimeSetting string         `gorm:"not null"`
	Character   string         `gorm:"not null"`
	Situation   string         `gorm:"not null"`
	Mood        string         `gorm:"not null"`
	StoryText   string         `gorm:"type:text;not null"`
	IsPublic    bool           `gorm:"default:false"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
} 