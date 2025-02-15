package handler

import (
	"library-service/internal/model"
	"library-service/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	fiber       *fiber.App
	userService port.UserService
	jwt         port.JWT
	validator   *validator.Validate
}

func NewRouteUserHandler(f *fiber.App, userService port.UserService, jwt port.JWT, validator *validator.Validate) *UserHandler {
	handler := &UserHandler{
		fiber:       f,
		userService: userService,
		jwt:         jwt,
		validator:   validator,
	}

	f.Post("/login", handler.Login)
	f.Get("/users", jwt.ValidateLibrarian, handler.GetAllMember)
	user := f.Group("/user")
	user.Post("/register", handler.RegisterHandler)

	librarian := f.Group("/librarian")
	librarian.Post("/register", handler.jwt.Validate)

	return handler
}

type RegisterBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
	Role     string `json:"role" validate:"required"`
}

func (u *UserHandler) GetAllMember(ctx *fiber.Ctx) error {
	users, err := u.userService.GetAllMember()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "get users success",
		"users":   users,
	})
}

func (u *UserHandler) RegisterHandler(ctx *fiber.Ctx) error {
	var input RegisterBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := u.validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	user := &model.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
		Role:     input.Role,
	}

	if err := u.userService.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "register success",
	})
}

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var input LoginBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	if err := u.validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := u.userService.Authentication(input.Username, input.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token":   token,
		"message": "login success",
	})
}
