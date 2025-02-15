package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:uuid;primary_key;"`
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	Username      string         `gorm:"uniqueIndex"`
	Password      string
	Name          string
	Role          string
	BorrowDetails []*BorrowDetail `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;"`
}

type BorrowDetail struct {
	gorm.Model
	BookName   string
	BorrowedAt int64
	ReturnedAt *int64
	User       *User `gorm:"foreignKey:UserID;references:ID"`
	UserID     string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
