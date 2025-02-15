package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"strings"
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

func (u *UserHandlerSuite) TestGetAllMember_Success() {
	u.mockService.EXPECT().GetAllMember().Return(u.mockUser, nil)

	u.fiber.Post("/users", u.handler.GetAllMember)

	req := httptest.NewRequest("POST", "/users", nil)
	resp, _ := u.fiber.Test(req)
	u.Equal(fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Equal("get users success", response["message"])
	u.NotNil(response["users"])
}

func (u *UserHandlerSuite) TestGetAllMember_Fail() {
	u.mockService.EXPECT().GetAllMember().Return(nil, fmt.Errorf(mock.Anything))

	u.fiber.Post("/users", u.handler.GetAllMember)

	req := httptest.NewRequest("POST", "/users", nil)
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Equal(mock.Anything, response["message"])
}

func (u *UserHandlerSuite) TestRegisterHandler_Success() {
	requestBody := `{
		"username": "testuser",
		"password": "password123",
		"name": "Test User",
		"role": "member"
	}`

	u.mockService.EXPECT().CreateUser(mock.Anything).Return(nil)

	req := httptest.NewRequest("POST", "/user/register", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusCreated, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Equal("register success", response["message"])
}

func (u *UserHandlerSuite) TestRegisterHandler_ValidationFail() {
	requestBody := `{
		"username": "testuser",
		"name": "Test User",
		"role": "member"
	}`

	req := httptest.NewRequest("POST", "/user/register", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Contains(response["error"], "Password")
}

func (u *UserHandlerSuite) TestRegisterHandler_CreateFailed() {
	requestBody := `{
		"username": "testuser",
		"password": "password123",
		"name": "Test User",
		"role": "member"
	}`

	u.mockService.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

	req := httptest.NewRequest("POST", "/user/register", bytes.NewBufferString(requestBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Equal(mock.Anything, response["message"])
}

func (u *UserHandlerSuite) TestLogin_Success() {
	inputBody := `{
		"username": "testuser",
		"password": "password123"
	}`

	u.mockService.EXPECT().Authentication(mock.Anything, mock.Anything).Return(mock.Anything, nil)

	req := httptest.NewRequest("POST", "/login", strings.NewReader(inputBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusOK, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Equal(mock.Anything, response["token"])
	u.Equal("login success", response["message"])
}

func (u *UserHandlerSuite) TestLogin_ValidationError() {
	inputBody := `{
		"username": "testuser"
	}`

	req := httptest.NewRequest("POST", "/login", strings.NewReader(inputBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Contains(response["error"].(string), "Password")
}

func (u *UserHandlerSuite) TestLogin_FailAuthentication() {
	inputBody := `{
		"username": "testuser",
		"password": "wrongpassword"
	}`

	u.mockService.On("Authentication", "testuser", "wrongpassword").Return("", fmt.Errorf("invalid username or password"))

	req := httptest.NewRequest("POST", "/login", strings.NewReader(inputBody))
	req.Header.Set("Content-Type", "application/json")
	resp, _ := u.fiber.Test(req)

	u.Equal(fiber.StatusBadRequest, resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var response map[string]interface{}
	_ = json.Unmarshal(body, &response)

	u.Equal("invalid username or password", response["message"])
}
