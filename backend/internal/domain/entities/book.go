package entities

import (
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID       string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Name     string
	Category string
	Status   string
	Stock    int
}

func (b *Book) BeforeCreate(tx *gorm.DB) (err error) {
	if b.ID == "" {
		b.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
