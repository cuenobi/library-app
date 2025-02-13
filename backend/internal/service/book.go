package service

import (
	"fmt"

	"library-service/internal/domain/entities"
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

func (b *BookService) CreateBook(book *entities.Book) error {
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

func (b *BookService) Borrow(bookID string) error {
	err := b.bookRepo.DecreaseBookStock(bookID)
	if err != nil {
		return err
	}

	return nil
}

func (b *BookService) Return(bookID string) error {
	err := b.bookRepo.IncreaseBookStock(bookID)
	if err != nil {
		return err
	}

	return nil
}
