package handler

import (
	"net/url"

	"library-service/internal/constant"
	"library-service/internal/model"
	"library-service/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

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

func (b *BookHandler) GetAllBook(ctx *fiber.Ctx) error {
	books, err := b.bookService.GetAllBook()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get books success",
		"books":   books,
	})
}

type CreateBookBody struct {
	Name     string `json:"name" validate:"required"`
	Category string `json:"category" validate:"required"`
	Stock    int    `json:"stock" validate:"required"`
}

func (b *BookHandler) CreateBook(ctx *fiber.Ctx) error {
	var input CreateBookBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := b.validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
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
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "create success",
	})
}

func (b *BookHandler) Borrow(ctx *fiber.Ctx) error {
	bookID, err := url.QueryUnescape(ctx.Params("book_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	userID, err := url.QueryUnescape(ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = b.bookService.Borrow(bookID, userID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}

func (b *BookHandler) Return(ctx *fiber.Ctx) error {
	bookID, err := url.QueryUnescape(ctx.Params("book_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	userID, err := url.QueryUnescape(ctx.Params("user_id"))
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	err = b.bookService.Return(bookID, userID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "success",
	})
}
