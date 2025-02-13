package entities

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID            string `gorm:"type:uuid;primary_key;default:uuid_generate_v4()"`
	Username      string
	Password      string
	Name          string
	Role          string
	BorrowDetails []*BorrowDetail
}

type BorrowDetail struct {
	gorm.Model
	BookName   string
	BorrowedAt *time.Time
	ReturnedAt *time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = strings.ReplaceAll(uuid.New().String(), "-", "")
	}
	return
}
