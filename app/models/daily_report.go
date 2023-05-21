package models

import (
	"time"
)

type DailyReport struct {
	ID               			uint 	   `gorm:"primaryKey"`
	UserId           			int
	User			 			User      `gorm:"foreignKey:UserId"`
	TodayDescription 			string    `gorm:"not null"`
	PreviousDayDescription      string    `gorm:"not null"`
	Blocker      				string    `gorm:"not null"`
	ProjectId        			int       `gorm:"not null"`
	CreatedAt        			time.Time `gorm:"not null"`
}