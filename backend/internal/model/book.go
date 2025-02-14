package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID        string `gorm:"type:uuid;primary_key;"`
	CreatedAt *time.Time
	UpdatedAt *time.Time
	Name      string `gorm:"uniqueIndex"`
	Category  string
	Status    string
	Stock     int
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
