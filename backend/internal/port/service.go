package port

import "library-service/internal/model"

type UserService interface {
	CreateUser(user *model.User) error
	CreateLibrarian(user *model.User) error
	Authentication(username, password string) (string, error)
}
