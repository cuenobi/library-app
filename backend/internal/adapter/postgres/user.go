package postgres

import (
	"library-service/internal/model"

	"gorm.io/gorm"
)

type User struct {
	conn *gorm.DB
}

func NewUser(conn *gorm.DB) *User {
	return &User{
		conn: conn,
	}
}

func (u *User) HasUsername(username string) (bool, error) {
	return false, nil
}

func (u *User) CreateUser(user *model.User) error {
	return nil
}

func (u *User) GetUserByUsername(username string) (*model.User, error) {
	return &model.User{}, nil
}
