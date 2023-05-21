package models

import (
	"time"
)

type User struct {
	ID               uint 	   	`gorm:"primaryKey"`
	Name             string    	`gorm:"type:varchar(255);not null"`
	Email            string    	`gorm:"uniqueIndex;not null"`
	Password         string    	`gorm:"not null"`
	ProjectId        int     	`gorm:"not null"`
	CreatedAt        time.Time 	`gorm:"not null"`
}

type SignUpInput struct {
	Name            string `json:"name" binding:"required"`
	Email           string `json:"email" binding:"required"`
	Password        string `json:"password" binding:"required,min=3"`
}

type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}