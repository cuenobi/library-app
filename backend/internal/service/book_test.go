package service

import (
	"testing"

	"library-service/internal/constant"
	"library-service/internal/domain/model"
	"library-service/mocks"

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

func (u *BookServiceSuite) SetupTest() {
	u.mockBookRepo = *mocks.NewBookRepository(u.T())
	u.mockJwt = *mocks.NewJWT(u.T())
	u.mockBook = &model.Book{
		Name:     "foo",
		Category: "scientific",
		Status:   constant.Available,
		Stock:    4,
	}
}

func (u *BookServiceSuite) TearDownTest() {
	u.mockBookRepo.AssertExpectations(u.T())
	u.mockJwt.AssertExpectations(u.T())
}

func (u *BookServiceSuite) SetupSubTest() {
	u.TearDownTest()
	u.SetupTest()
}
