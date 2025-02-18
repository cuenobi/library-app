package handler

import (
	"library-service/internal/model"
	"library-service/internal/port"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// Define response structs for Swagger documentation

// UserResponse is the response format for GetAllMember, RegisterHandler, and Login
type UserResponse struct {
	Message string      `json:"message"`
	Token   string      `json:"token,omitempty"`
	Role    string      `json:"role,omitempty"`
	Users   interface{} `json:"users,omitempty"` // You can replace interface{} with a specific type if possible
}

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

// GetAllMember godoc
// @Summary Get all members
// @Description Retrieve all registered users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /users [get]
func (u *UserHandler) GetAllMember(ctx *fiber.Ctx) error {
	users, err := u.userService.GetAllMember()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(UserResponse{
		Message: "get users success",
		Users:   users,
	})
}

// RegisterHandler godoc
// @Summary Register a new user
// @Description Create a new user account
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body RegisterBody true "User registration details"
// @Success 201 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /user/register [post]
func (u *UserHandler) RegisterHandler(ctx *fiber.Ctx) error {
	var input RegisterBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := u.validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	user := &model.User{
		Username: input.Username,
		Password: input.Password,
		Name:     input.Name,
		Role:     input.Role,
	}

	if err := u.userService.CreateUser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(UserResponse{
		Message: "register success",
	})
}

type LoginBody struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

// Login godoc
// @Summary User login
// @Description Authenticate user and return JWT token
// @Tags users
// @Accept  json
// @Produce  json
// @Param request body LoginBody true "User login details"
// @Success 200 {object} UserResponse
// @Failure 400 {object} ErrorResponse
// @Router /login [post]
func (u *UserHandler) Login(ctx *fiber.Ctx) error {
	var input LoginBody
	if err := ctx.BodyParser(&input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	if err := u.validator.Struct(input); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	token, role, err := u.userService.Authentication(input.Username, input.Password)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(ErrorResponse{
			Error: err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(UserResponse{
		Token:   token,
		Role:    role,
		Message: "login success",
	})
}
