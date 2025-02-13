package port

import "library-service/internal/domain/entities"

type UserRepository interface {
	HasUsername(username string) (bool, error)
	CreateUser(user *entities.User) error
	GetUserByUsername(username string) (*entities.User, error)
}

type BookRepository interface {
	HasBookName(name string) (bool, error)
	CreateBook(book *entities.Book) error
	GetBookByID(id string) (*entities.Book, error)
	DecreaseBookStock(id string) error
	IncreaseBookStock(id string) error
}
