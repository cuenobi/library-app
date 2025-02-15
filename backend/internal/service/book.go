package service

import (
	"fmt"

	"library-service/internal/model"
	"library-service/internal/port"
)

type BookService struct {
	bookRepo port.BookRepository
}

func NewBookService(bookRepo port.BookRepository) *BookService {
	return &BookService{
		bookRepo: bookRepo,
	}
}

func (b *BookService) CreateBook(book *model.Book) error {
	bookExist, err := b.bookRepo.HasBookName(book.Name)
	if err != nil {
		return err
	}
	if bookExist {
		return fmt.Errorf("book name already exist")
	}

	err = b.bookRepo.CreateBook(book)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookService) Borrow(bookID, userID string) error {
	err := b.bookRepo.DecreaseBookStockAndAddUpdateBorrowDetail(bookID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookService) Return(bookID, userID string) error {
	err := b.bookRepo.IncreaseBookStockAndUpdateBorrowDetail(bookID, userID)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookService) GetAllBook() ([]*model.Book, error) {
	books, err := b.bookRepo.GetAllBook()
	if err != nil {
		return nil, err
	}

	return books, nil
}
