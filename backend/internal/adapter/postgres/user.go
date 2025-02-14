package postgres

import "gorm.io/gorm"

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
