package model

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	Title         string         `gorm:"not null" json:"title"`
	Author        string         `gorm:"not null" json:"author"`
	Genre         string         `json:"genre"`
	PublishedDate time.Time      `json:"published_date"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}
