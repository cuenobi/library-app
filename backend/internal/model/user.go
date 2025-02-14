package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	Username      string `gorm:"uniqueIndex"`
	Password      string
	Name          string
	Role          string
	BorrowDetails []*BorrowDetail `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
	DeletedAt     gorm.DeletedAt  `gorm:"index"`
}

type BorrowDetail struct {
	gorm.Model
	BookName   string
	BorrowedAt *time.Time
	ReturnedAt *time.Time
	User       User `gorm:"foreignKey:UserID;references:ID"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
