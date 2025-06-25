package models

import (
	"time"
)

type Favorite struct {
	ID          uint      `gorm:"primaryKey"`
	UserID      uint      `gorm:"not null"`
	User        User      `gorm:"foreignKey:UserID"`
	SituationID uint      `gorm:"not null"`
	Situation   Situation `gorm:"foreignKey:SituationID"`
	CreatedAt   time.Time
} 