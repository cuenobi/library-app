package handler

import (
	"testing"
	"time"

	"library-service/internal/constant"
	"library-service/internal/model"
	"library-service/mocks"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerSuite))
}

type UserHandlerSuite struct {
	suite.Suite
	handler     *UserHandler
	fiber       *fiber.App
	mockService mocks.UserService
	mockJwt     mocks.JWT
	mockUser    []*model.User
}

func (u *UserHandlerSuite) SetupTest() {
	u.mockService = *mocks.NewUserService(u.T())
	u.mockJwt = *mocks.NewJWT(u.T())
	u.fiber = fiber.New()
	u.handler = NewRouteUserHandler(u.fiber, &u.mockService, &u.mockJwt, validator.New())

	u.mockUser = []*model.User{
		{
			ID:        mock.Anything,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			UpdatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			Name:      mock.Anything,
			Role:      constant.MemberRole,
		},
		{
			ID:        mock.Anything,
			CreatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			UpdatedAt: func(t time.Time) *time.Time { return &t }(time.Now()),
			Name:      mock.Anything,
			Role:      constant.Librarian,
		},
	}
}

func (u *UserHandlerSuite) TearDownTest() {
	u.mockService.AssertExpectations(u.T())
	u.mockJwt.AssertExpectations(u.T())
}

func (u *UserHandlerSuite) SetupSubTest() {
	u.TearDownTest()
	u.SetupTest()
}

// func (u *UserHandlerSuite) Test