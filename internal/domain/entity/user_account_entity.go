package entity

import (
	"time"
)

type UserAccount struct {
	ID           int `gorm:"primaryKey;autoIncrement:true"`
	Name         string
	DayOfBirth   string
	Gender       string
	PasswordHash string
	Address      string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
