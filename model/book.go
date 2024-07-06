package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title         string `gorm:"not null"`
	Author        string `gorm:"not null"`
	Genre         string
	PublishedDate time.Time
}
