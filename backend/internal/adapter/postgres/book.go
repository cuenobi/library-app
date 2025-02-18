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

func (b *Book) GetAllBook() ([]*model.Book, error) {
	var books []*model.Book
	err := b.db.Model(&model.Book{}).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
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
	result := b.db.First(&book, id)
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
		BookID:     bookID,
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

	var borrowDetail model.BorrowDetail
	err = tx.First(&borrowDetail, "book_id = ? AND user_id = ? AND returned_at IS NULL", book.ID, userID).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Model(borrowDetail).
		Where("book_id = ? AND user_id = ?", book.ID, userID).
		Update("returned_at", time.Now().Unix()).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}
