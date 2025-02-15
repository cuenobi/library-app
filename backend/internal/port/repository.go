package port

import "library-service/internal/model"

type UserRepository interface {
	GetAllMember() ([]*model.User, error)
	HasUsername(username string) (bool, error)
	CreateUser(user *model.User) error
	GetUserByUsername(username string) (*model.User, error)
}

type BookRepository interface {
	GetAllBook() ([]*model.Book, error)
	HasBookName(name string) (bool, error)
	CreateBook(book *model.Book) error
	GetBookByID(id string) (*model.Book, error)
	DecreaseBookStockAndAddUpdateBorrowDetail(bookID, userID string) error
	IncreaseBookStockAndUpdateBorrowDetail(bookID, userID string) error
}
