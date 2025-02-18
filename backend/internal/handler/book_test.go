package handler

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"library-service/internal/model"
	"library-service/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestBookHandlerSuite(t *testing.T) {
	suite.Run(t, new(BookHandlerSuite))
}

type BookHandlerSuite struct {
	suite.Suite
	handler     *BookHandler
	fiber       *fiber.App
	mockService mocks.BookService
	mockJwt     mocks.JWT
	mockBook    []*model.Book
}

func (b *BookHandlerSuite) SetupTest() {
	b.mockService = *mocks.NewBookService(b.T())
	b.mockJwt = *mocks.NewJWT(b.T())
	b.fiber = fiber.New()
	b.handler = NewRouteBookHandler(b.fiber, &b.mockService, &b.mockJwt, validator.New())

	b.mockBook = []*model.Book{
		{
			ID:        mock.Anything,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			UpdatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			Name:      mock.Anything,
			Category:  mock.Anything,
			Status:    mock.Anything,
			Stock:     3,
		},
		{
			ID:        mock.Anything,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			UpdatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			Name:      mock.Anything,
			Category:  mock.Anything,
			Status:    mock.Anything,
			Stock:     2,
		},
	}
}

func (b *BookHandlerSuite) TearDownTest() {
	b.mockService.AssertExpectations(b.T())
	b.mockJwt.AssertExpectations(b.T())
}

func (b *BookHandlerSuite) SetupSubTest() {
	b.TearDownTest()
	b.SetupTest()
}

func (b *BookHandlerSuite) TestGetAllBookSuccess() {
	// Mocking the service method to return mockBooks with no error
	b.mockService.EXPECT().GetAllBook().Return(b.mockBook, nil).Once()

	// Register route handler in fiber app
	b.fiber.Post("/books", b.handler.GetAllBook)

	// Simulate an HTTP request to the /books route
	req := httptest.NewRequest("POST", "/books", nil)
	resp, err := b.fiber.Test(req)

	// Check if there is no error
	b.Require().NoError(err)

	// Check the status code is 200 OK
	b.Assert().Equal(resp.StatusCode, fiber.StatusOK)

	// Check if the response body is as expected
	var result fiber.Map
	err = json.NewDecoder(resp.Body).Decode(&result)
	b.Require().NoError(err)

	// Check the response structure
	b.Assert().Equal(result["message"], "get books success")
	b.Assert().Len(result["books"].([]interface{}), len(b.mockBook))
}

// func (b *BookHandlerSuite) TestGetAllBookError() {
// 	// Mocking the service method to return an error
// 	b.mockService.EXPECT().GetAllBook().Return(nil, fmt.Errorf("some error")).Once()

// 	// Register route handler in fiber app
// 	b.fiber.Post("/books", b.handler.GetAllBook)

// 	// Simulate an HTTP request to the /books route
// 	req := httptest.NewRequest("POST", "/books", nil)
// 	resp, err := b.fiber.Test(req)

// 	// Check if there is no error
// 	b.Require().NoError(err)

// 	// Check the status code is 400 Bad Request
// 	b.Assert().Equal(resp.StatusCode, fiber.StatusBadRequest)

// 	// Check if the response body contains the error message
// 	var result fiber.Map
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	b.Require().NoError(err)

// 	// Assert the error message
// 	b.Assert().Equal(result["message"], "some error")
// }

func (b *BookHandlerSuite) TestCreateBookSuccess() {
	// Mock the bookService's CreateBook method to return no error
	b.mockService.EXPECT().CreateBook(mock.Anything).Return(nil).Once()

	// Register route handler in fiber app
	b.fiber.Post("/books", b.handler.CreateBook)

	// Create request body for creating a new book
	requestBody := `{
		"name": "Go Programming",
		"category": "Programming",
		"stock": 10
	}`

	// Simulate an HTTP request to the /books route
	req := httptest.NewRequest("POST", "/books", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := b.fiber.Test(req)

	// Check if there is no error
	b.Require().NoError(err)

	// Check the status code is 201 Created
	b.Assert().Equal(resp.StatusCode, fiber.StatusCreated)

	// Check if the response message is as expected
	var result fiber.Map
	err = json.NewDecoder(resp.Body).Decode(&result)
	b.Require().NoError(err)
	b.Assert().Equal(result["message"], "create success")
}

func (b *BookHandlerSuite) TestCreateBookValidationError() {
	// Register route handler in fiber app
	b.fiber.Post("/books", b.handler.CreateBook)

	// Simulate an HTTP request with invalid data (missing required fields)
	requestBody := `{
		"name": "Go Programming",
		"category": "",
		"stock": 10
	}`

	// Simulate an HTTP request to the /books route
	req := httptest.NewRequest("POST", "/books", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := b.fiber.Test(req)

	// Check if there is no error
	b.Require().NoError(err)

	// Check the status code is 400 Bad Request
	b.Assert().Equal(fiber.StatusBadRequest, resp.StatusCode)

	// Check if the response contains validation error
	var result fiber.Map
	err = json.NewDecoder(resp.Body).Decode(&result)
	b.Require().NoError(err)

	// Assert that error contains validation error message
	b.Assert().Contains(result["error"].(string), "Category")
}

func (b *BookHandlerSuite) TestCreateBookServiceError() {
	// Mock the bookService's CreateBook method to return an error
	b.mockService.EXPECT().CreateBook(mock.Anything).Return(fmt.Errorf("unable to create book")).Once()

	// Register route handler in fiber app
	b.fiber.Post("/books", b.handler.CreateBook)

	// Create valid request body for creating a new book
	requestBody := `{
		"name": "Go Programming",
		"category": "Programming",
		"stock": -1
	}`

	// Simulate an HTTP request to the /books route
	req := httptest.NewRequest("POST", "/books", strings.NewReader(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := b.fiber.Test(req)

	// Check if there is no error
	b.Require().NoError(err)

	// Check the status code is 400 Bad Request
	b.Assert().Equal(resp.StatusCode, fiber.StatusBadRequest)

	// Check if the response contains the error message
	var result fiber.Map
	err = json.NewDecoder(resp.Body).Decode(&result)
	b.Require().NoError(err)

	// Assert that error contains service error message
	b.Assert().Equal(result["error"], "unable to create book")
}

// func (b *BookHandlerSuite) TestBorrowSuccess() {
// 	// Mocking the service method to return no error
// 	bookID := "123"
// 	userID := "456"
// 	b.mockService.EXPECT().Borrow(bookID, userID).Return(nil).Once()

// 	// Register route handler in fiber app
// 	b.fiber.Get("/book/borrow/:user_id/:book_id", b.handler.Borrow)

// 	// Simulate an HTTP request to borrow a book
// 	req := httptest.NewRequest("GET", "/book/borrow/456/123", nil)
// 	resp, err := b.fiber.Test(req)

// 	// Check if there is no error
// 	b.Require().NoError(err)

// 	// Check the status code is 200 OK
// 	b.Assert().Equal(fiber.StatusOK, resp.StatusCode)

// 	// Check if the response body is as expected
// 	var result fiber.Map
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	b.Require().NoError(err)

// 	// Validate the response message
// 	b.Assert().Equal("success", result["message"])
// }

// func (b *BookHandlerSuite) TestBorrowFailServiceError() {
// 	// Mocking the service method to return an error
// 	bookID := "123"
// 	userID := "456"
// 	b.mockService.EXPECT().Borrow(bookID, userID).Return(fmt.Errorf("borrow failed")).Once()

// 	// Register route handler in fiber app
// 	b.fiber.Get("/book/borrow/:user_id/:book_id", b.handler.Borrow)

// 	// Simulate an HTTP request to borrow a book
// 	req := httptest.NewRequest("GET", "/book/borrow/456/123", nil)
// 	resp, err := b.fiber.Test(req)

// 	// Check if there is no error
// 	b.Require().NoError(err)

// 	// Check the status code is 400 Bad Request
// 	b.Assert().Equal(fiber.StatusBadRequest, resp.StatusCode)

// 	// Check if the response body contains the expected error message
// 	var result string
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	b.Require().NoError(err)

// 	b.Assert().Equal("borrow failed", result)
// }

// func (b *BookHandlerSuite) TestReturnSuccess() {
// 	// Mocking the service method to return no error
// 	bookID := "123"
// 	userID := "456"
// 	b.mockService.EXPECT().Return(bookID, userID).Return(nil).Once()

// 	// Register route handler in fiber app
// 	b.fiber.Get("/book/return/:user_id/:book_id", b.handler.Return)

// 	// Simulate an HTTP request to return a book
// 	req := httptest.NewRequest("GET", "/book/return/456/123", nil)
// 	resp, err := b.fiber.Test(req)

// 	// Check if there is no error
// 	b.Require().NoError(err)

// 	// Check the status code is 200 OK
// 	b.Assert().Equal(fiber.StatusOK, resp.StatusCode)

// 	// Check if the response body is as expected
// 	var result fiber.Map
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	b.Require().NoError(err)

// 	// Validate the response message
// 	b.Assert().Equal("success", result["message"])
// }

// func (b *BookHandlerSuite) TestReturnFailServiceError() {
// 	// Mocking the service method to return an error
// 	bookID := "123"
// 	userID := "456"
// 	b.mockService.EXPECT().Return(bookID, userID).Return(fmt.Errorf("return failed")).Once()

// 	// Register route handler in fiber app
// 	b.fiber.Get("/book/return/:user_id/:book_id", b.handler.Return)

// 	// Simulate an HTTP request to return a book
// 	req := httptest.NewRequest("GET", "/book/return/456/123", nil)
// 	resp, err := b.fiber.Test(req)

// 	// Check if there is no error
// 	b.Require().NoError(err)

// 	// Check the status code is 400 Bad Request
// 	b.Assert().Equal(fiber.StatusBadRequest, resp.StatusCode)

// 	// Check if the response body contains the expected error message
// 	var result string
// 	err = json.NewDecoder(resp.Body).Decode(&result)
// 	b.Require().NoError(err)

// 	// Assert that the error message matches the service error
// 	b.Assert().Equal("return failed", result)
// }
