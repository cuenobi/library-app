package postgres

import (
	"time"

	"library-service/internal/model"

	"gorm.io/gorm"
)

type Book struct {
	db *gorm.DB
}

func NewBook(db *gorm.DB) *Book {
	return &Book{
		db: db,
	}
}

func (b *Book) HasBookName(name string) (bool, error) {
	var count int64
	b.db.Model(&model.Book{}).Where("name = ?", name).Count(&count)
	return count > 0, nil
}

func (b *Book) CreateBook(book *model.Book) error {
	result := b.db.Create(book)
	return result.Error
}

func (b *Book) GetBookByID(id string) (*model.Book, error) {
	var book *model.Book
	result := b.db.First(&book, "10")
	return book, result.Error
}

func (b *Book) DecreaseBookStockAndAddUpdateBorrowDetail(bookID, userID string) error {
	tx := b.db.Begin()

	var book model.Book
	err := tx.Model(&model.Book{}).
		Where("id = ?", bookID).
		Update("stock", gorm.Expr("GREATEST(stock - ?, 0)", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.First(&book, "id = ?", bookID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Create(&model.BorrowDetail{
		BookName:   book.Name,
		BorrowedAt: time.Now().Unix(),
		UserID:     userID,
	}).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (b *Book) IncreaseBookStockAndUpdateBorrowDetail(bookID, userID string) error {
	tx := b.db.Begin()

	var book model.Book
	err := tx.Model(&model.Book{}).
		Where("id = ?", bookID).
		Update("stock", gorm.Expr("stock + ?", 1)).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.First(&book, "id = ?", bookID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(&model.BorrowDetail{}).
		Where("name = ?", book.Name).
		Update("returned_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
