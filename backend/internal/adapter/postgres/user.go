package postgres

import (
	"library-service/internal/model"

	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (u *User) HasUsername(username string) (bool, error) {
	var count int64
	u.db.Model(&model.User{}).Where("username = ?", username).Count(&count)
	return count > 0, nil
}

func (u *User) CreateUser(user *model.User) error {
	result := u.db.Create(user)
	return result.Error
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	var user *model.User
	err := u.db.First(&user, "username = ?", username).Error
	return user, err
}
