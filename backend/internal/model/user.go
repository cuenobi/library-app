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
	ID         string `gorm:"type:uuid"`
	BookName   string
	BorrowedAt int64
	ReturnedAt *int64
	UserID     string
	BookID     string
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}

func (u *BorrowDetail) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
