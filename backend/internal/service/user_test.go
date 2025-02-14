package service

import (
	"fmt"
	"testing"

	"library-service/internal/constant"
	"library-service/internal/model"
	"library-service/mocks"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
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
	mockPassword  string
}

func (u *UserServiceSuite) SetupTest() {
	u.mockUserRepo = *mocks.NewUserRepository(u.T())
	u.mockJwt = *mocks.NewJWT(u.T())

	u.mockPassword = "password"
	hashPwd, _ := bcrypt.GenerateFromPassword([]byte(u.mockPassword), bcrypt.DefaultCost)

	u.mockUser = &model.User{
		Username: "foo",
		Password: string(hashPwd),
		Name:     "bar",
		Role:     constant.MemberRole,
	}

	u.mockLibrarian = &model.User{
		Username: "bar",
		Password: string(hashPwd),
		Name:     "foo",
		Role:     constant.Librarian,
	}

	u.service = NewUserService(&u.mockUserRepo, &u.mockJwt)
}

func (u *UserServiceSuite) TearDownTest() {
	u.mockUserRepo.AssertExpectations(u.T())
	u.mockJwt.AssertExpectations(u.T())
}

func (u *UserServiceSuite) SetupSubTest() {
	u.TearDownTest()
	u.SetupTest()
}

func (u *UserServiceSuite) TestCreateUserSuccess() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(nil)

	err := u.service.CreateUser(u.mockUser)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestCreateUserInvalidRole() {
	err := u.service.CreateUser(u.mockLibrarian)
	u.Require().Error(err)
	u.Require().Equal("invalid role", err.Error())
}

func (u *UserServiceSuite) TestCreateUserUsernameAlreadyExist() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(true, nil)

	err := u.service.CreateUser(u.mockUser)
	u.Require().Error(err)
	u.Require().Equal("username already exist", err.Error())
}

func (u *UserServiceSuite) TestCreateUserHasUsernameError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

	err := u.service.CreateUser(u.mockUser)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestCreateUserCreateUserError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := u.service.CreateUser(u.mockUser)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestCreateLibrarianSuccess() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(nil)

	err := u.service.CreateLibrarian(u.mockLibrarian)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestCreateLibrarianInvalidRole() {
	err := u.service.CreateLibrarian(u.mockUser)
	u.Require().Error(err)
	u.Require().Equal("invalid role", err.Error())
}

func (u *UserServiceSuite) TestCreateLibrarianUsernameAlreadyExist() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(true, nil)

	err := u.service.CreateLibrarian(u.mockLibrarian)
	u.Require().Error(err)
	u.Require().Equal("username already exist", err.Error())
}

func (u *UserServiceSuite) TestCreateLibrarianHasUsernameError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, fmt.Errorf(mock.Anything))

	err := u.service.CreateLibrarian(u.mockLibrarian)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestCreateLibrarianCreateUserError() {
	u.mockUserRepo.EXPECT().HasUsername(mock.Anything).Return(false, nil)
	u.mockUserRepo.EXPECT().CreateUser(mock.Anything).Return(fmt.Errorf(mock.Anything))

	err := u.service.CreateLibrarian(u.mockLibrarian)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestAuthenticationSuccess() {
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)
	u.mockJwt.EXPECT().Generate(mock.Anything, mock.Anything).Return(mock.Anything)

	_, err := u.service.Authentication(u.mockUser.Username, u.mockPassword)
	u.Require().NoError(err)
}

func (u *UserServiceSuite) TestAuthenticationGetUserError() {
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(nil, fmt.Errorf(mock.Anything))

	_, err := u.service.Authentication(u.mockUser.Username, u.mockPassword)
	u.Require().Error(err)
	u.Require().Equal(mock.Anything, err.Error())
}

func (u *UserServiceSuite) TestAuthenticationCompareError() {
	u.mockUserRepo.EXPECT().GetUserByUsername(mock.Anything).Return(u.mockUser, nil)

	_, err := u.service.Authentication(u.mockUser.Username, u.mockUser.Password)
	u.Require().Error(err)
}