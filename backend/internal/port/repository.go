package port

import "library-service/internal/model"

type UserRepository interface {
	HasUsername(username string) (bool, error)
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
}

type BookRepository interface {
	HasBookName(name string) (bool, error)
	CreateBook(book *model.Book) error
	GetBookByID(id string) (*model.Book, error)
	DecreaseBookStock(id string) error
	IncreaseBookStock(id string) error
}
