package service

import (
	"fmt"
	"testing"

	"library-service/internal/constant"
	"library-service/internal/model"
	"library-service/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestBookServiceSuite(t *testing.T) {
	suite.Run(t, new(BookServiceSuite))
}

type BookServiceSuite struct {
	suite.Suite
	service      *BookService
	mockBookRepo mocks.BookRepository
	mockJwt      mocks.JWT
	mockBook     *model.Book
}

func (b *BookServiceSuite) SetupTest() {
	b.mockBookRepo = *mocks.NewBookRepository(b.T())
	b.mockJwt = *mocks.NewJWT(b.T())
	b.mockBook = &model.Book{
		Name:     "foo",
		Category: "scientific",
		Status:   constant.Available,
		Stock:    4,
	}
	b.service = NewBookService(&b.mockBookRepo)
}

func (b *BookServiceSuite) TearDownTest() {
	b.mockBookRepo.AssertExpectations(b.T())
	b.mockJwt.AssertExpectations(b.T())
}

func (b *BookServiceSuite) SetupSubTest() {
	b.TearDownTest()
	b.SetupTest()
}

func (b *BookServiceSuite) TestCreateBookSuccess() {
	b.mockBookRepo.EXPECT().HasBookName(mock.Anything).Return(false, nil)
	b.mockBookRepo.EXPECT().CreateBook(mock.Anything).Return(nil)

	err := b.service.CreateBook(b.mockBook)
	b.Require().NoError(err)
}

func (b *BookServiceSuite) TestCreateBookHasBookNameError() {
	b.mockBookRepo.EXPECT().HasBookName(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

	err := b.service.CreateBook(b.mockBook)
	b.Require().Error(err)
	b.Require().Equal(mock.Anything, err.Error())
}

func (b *BookServiceSuite) TestCreateBookAlreadyBookName() {
	b.mockBookRepo.EXPECT().HasBookName(mock.Anything).Return(true, nil)

	err := b.service.CreateBook(b.mockBook)
	b.Require().Error(err)
	b.Require().Equal("book name already exist", err.Error())
}

func (b *BookServiceSuite) TestCreateBookError() {
	b.mockBookRepo.EXPECT().HasBookName(mock.Anything).Return(false, nil)
	b.mockBookRepo.EXPECT().CreateBook(mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := b.service.CreateBook(b.mockBook)
	b.Require().Error(err)
	b.Require().Equal(mock.Anything, err.Error())
}

func (b *BookServiceSuite) TestBorrowSuccess() {
	b.mockBookRepo.EXPECT().DecreaseBookStockAndAddUpdateBorrowDetail(mock.Anything, mock.Anything).Return(nil)

	err := b.service.Borrow("mock", mock.Anything)
	b.Require().NoError(err)
}

func (b *BookServiceSuite) TestBorrowError() {
	b.mockBookRepo.EXPECT().DecreaseBookStockAndAddUpdateBorrowDetail(mock.Anything, mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := b.service.Borrow("mock", mock.Anything)
	b.Require().Error(err)
	b.Require().Equal(mock.Anything, err.Error())
}

func (b *BookServiceSuite) TestReturnSuccess() {
	b.mockBookRepo.EXPECT().IncreaseBookStockAndUpdateBorrowDetail(mock.Anything, mock.Anything).Return(nil)

	err := b.service.Return("mock", mock.Anything)
	b.Require().NoError(err)
}

func (b *BookServiceSuite) TestReturnError() {
	b.mockBookRepo.EXPECT().IncreaseBookStockAndUpdateBorrowDetail(mock.Anything, mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := b.service.Return("mock", mock.Anything)
	b.Require().Error(err)
	b.Require().Equal(mock.Anything, err.Error())
}
