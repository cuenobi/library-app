package handler

import (
	"library-service/internal/port"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	fiber       *fiber.App
	userService port.UserService
	jwt         port.JWT
}

func NewRouteUserHandler(f *fiber.App, userService port.UserService, jwt port.JWT) {
	handler := &UserHandler{
		fiber:       f,
		userService: userService,
		jwt:         jwt,
	}

	user := f.Group("/user")
	user.Post("/register", handler.RegisterHandler)

	librarian := f.Group("/librarian")
	librarian.Post("/register", handler.jwt.Validate)
}

func (u *UserHandler) RegisterHandler(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
	})
}
