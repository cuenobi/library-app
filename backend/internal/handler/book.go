package handler

import (
	"net/url"

	"library-service/internal/constant"
	"library-service/internal/model"
	"library-service/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Define response structs for Swagger documentation

// BookResponse is the response format for GetAllBook
type BookResponse struct {
	Message string      `json:"message"`
	Books   interface{} `json:"books,omitempty"` // You can replace interface{} with a specific type if possible
}

// ErrorResponse is a generic response format for errors
type ErrorResponse struct {
	Error string `json:"error"`
}

type BookHandler struct {
	fiber       *fiber.App
	bookService port.BookService
	jwt         port.JWT
	validator   *validator.Validate
}

func NewRouteBookHandler(f *fiber.App, bookService port.BookService, jwt port.JWT, validator *validator.Validate) *BookHandler {
	handler := &BookHandler{
		fiber:       f,
		bookService: bookService,
		jwt:         jwt,
		validator:   validator,
	}

	f.Get("/books", jwt.ValidateLibrarian, handler.GetAllBook)
	book := f.Group("/book")
	book.Post("/create", jwt.ValidateLibrarian, handler.CreateBook)
	book.Post("/borrow/:user_id/:book_id", jwt.ValidateLibrarian, handler.Borrow)
	book.Post("/return/:user_id/:book_id", jwt.ValidateLibrarian, handler.Return)

	return handler
}

// GetAllBook godoc
// @Summary Get all books
// @Description Retrieve all books in the library
// @Tags books
// @Accept  json
// @Produce  json
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Router /books [get]
func (b *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	books, err := b.bookService.GetAllBook()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	response := BookResponse{
		Message: "get books success",
		Books:   books,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

type CreateBookBody struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	Stock    int    `json:"stock" validate:"required"`
}

// CreateBook godoc
// @Summary Create a new book
// @Description Add a new book to the library
// @Tags books
// @Accept  json
// @Produce  json
// @Param request body CreateBookBody true "Book details"
// @Success 201 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Router /book/create [post]
func (b *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	var input CreateBookBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := b.validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	book := &model.Book{
		Name:     input.Name,
		Category: input.Category,
		Stock:    input.Stock,
		Status: func() string {
			if input.Stock > 0 {
				return constant.Available
			}
			return constant.OutOfStock
		}(),
	}

	err := b.bookService.CreateBook(book)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	// Create success response
	return ctx.Status(fiber.StatusCreated).JSON(BookResponse{
		Message: "create success",
	})
}

// Borrow godoc
// @Summary Borrow a book
// @Description Borrow a book by user ID and book ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param user_id path string true "User ID"
// @Param book_id path string true "Book ID"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Router /book/borrow/{user_id}/{book_id} [post]
func (b *BookHandler) Borrow(ctx *fiber.Ctx) error {
	bookID, err := url.QueryUnescape(ctx.Params("book_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	userID, err := url.QueryUnescape(ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	err = b.bookService.Borrow(bookID, userID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(BookResponse{
		Message: "borrow success",
	})
}

// Return godoc
// @Summary Return a book
// @Description Return a borrowed book by user ID and book ID
// @Tags books
// @Accept  json
// @Produce  json
// @Param user_id path string true "User ID"
// @Param book_id path string true "Book ID"
// @Success 200 {object} BookResponse
// @Failure 400 {object} ErrorResponse
// @Router /book/return/{user_id}/{book_id} [post]
func (b *BookHandler) Return(ctx *fiber.Ctx) error {
	bookID, err := url.QueryUnescape(ctx.Params("book_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	userID, err := url.QueryUnescape(ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	err = b.bookService.Return(bookID, userID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(BookResponse{
		Message: "return success",
	})
}