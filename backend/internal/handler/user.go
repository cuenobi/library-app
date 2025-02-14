package handler

import (
	"library-service/internal/model"
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

type RegisterBody struct {
	Username *string `json:"username"`
	Password *string `json:"password"`
	Name     *string `json:"name"`
	Role     *string `json:"role"`
}

func (u *UserHandler) RegisterHandler(ctx *fiber.Ctx) error {
	var input RegisterBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	user := &model.User{
		Username: *input.Username,
		Password: *input.Password,
		Name:     *input.Name,
		Role:     *input.Role,
	}

	if err := u.userService.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
	})
}
