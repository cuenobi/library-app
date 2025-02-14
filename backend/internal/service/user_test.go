package service

import (
	"testing"

	"library-service/internal/constant"
	"library-service/internal/domain/model"
	"library-service/mocks"

	"github.com/stretchr/testify/suite"
)

func TestUserServiceSuite(t *testing.T) {
	suite.Run(t, new(UserServiceSuite))
}

type UserServiceSuite struct {
	suite.Suite
	service       *UserService
	mockUserRepo  mocks.UserRepository
	mockJwt       mocks.JWT
	mockUser      *model.User
	mockLibrarian *model.User
}

func (u *UserServiceSuite) SetupTest() {
	u.mockUserRepo = *mocks.NewUserRepository(u.T())
	u.mockJwt = *mocks.NewJWT(u.T())
	u.mockUser = &model.User{
		Username: "foo",
		Password: "password",
		Name:     "bar",
		Role:     constant.MemberRole,
	}
	u.mockLibrarian = &model.User{
		Username: "bar",
		Password: "password",
		Name:     "foo",
		Role:     constant.Librarian,
	}
}

func (u *UserServiceSuite) TearDownTest() {
	u.mockUserRepo.AssertExpectations(u.T())
	u.mockJwt.AssertExpectations(u.T())
}

func (u *UserServiceSuite) SetupSubTest() {
	u.TearDownTest()
	u.SetupTest()
}