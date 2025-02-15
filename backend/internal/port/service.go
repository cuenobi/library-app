package port

import "library-service/internal/model"

type UserService interface {
	GetAllMember() ([]*model.User, error)
	CreateUser(user *model.User) error
	CreateLibrarian(user *model.User) error
	Authentication(username, password string) (string, error)
}

type BookService interface {
	GetAllBook() ([]*model.Book, error)
	CreateBook(book *model.Book) error
	Borrow(bookID, userID string) error
	Return(bookID, userID string) error
}
