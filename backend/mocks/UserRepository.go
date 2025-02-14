// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	model "library-service/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// UserRepository is an autogenerated mock type for the UserRepository type
type UserRepository struct {
	mock.Mock
}

type UserRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *UserRepository) EXPECT() *UserRepository_Expecter {
	return &UserRepository_Expecter{mock: &_m.Mock}
}

// CreateUser provides a mock function with given fields: user
func (_m *UserRepository) CreateUser(user *model.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for CreateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UserRepository_CreateUser_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateUser'
type UserRepository_CreateUser_Call struct {
	*mock.Call
}

// CreateUser is a helper method to define mock.On call
//   - user *model.User
func (_e *UserRepository_Expecter) CreateUser(user interface{}) *UserRepository_CreateUser_Call {
	return &UserRepository_CreateUser_Call{Call: _e.mock.On("CreateUser", user)}
}

func (_c *UserRepository_CreateUser_Call) Run(run func(user *model.User)) *UserRepository_CreateUser_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.User))
	})
	return _c
}

func (_c *UserRepository_CreateUser_Call) Return(_a0 error) *UserRepository_CreateUser_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *UserRepository_CreateUser_Call) RunAndReturn(run func(*model.User) error) *UserRepository_CreateUser_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *UserRepository) GetUserByUsername(username string) (*model.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *model.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *model.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_GetUserByUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserByUsername'
type UserRepository_GetUserByUsername_Call struct {
	*mock.Call
}

// GetUserByUsername is a helper method to define mock.On call
//   - username string
func (_e *UserRepository_Expecter) GetUserByUsername(username interface{}) *UserRepository_GetUserByUsername_Call {
	return &UserRepository_GetUserByUsername_Call{Call: _e.mock.On("GetUserByUsername", username)}
}

func (_c *UserRepository_GetUserByUsername_Call) Run(run func(username string)) *UserRepository_GetUserByUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserRepository_GetUserByUsername_Call) Return(_a0 *model.User, _a1 error) *UserRepository_GetUserByUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_GetUserByUsername_Call) RunAndReturn(run func(string) (*model.User, error)) *UserRepository_GetUserByUsername_Call {
	_c.Call.Return(run)
	return _c
}

// HasUsername provides a mock function with given fields: username
func (_m *UserRepository) HasUsername(username string) (bool, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for HasUsername")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(username)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UserRepository_HasUsername_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasUsername'
type UserRepository_HasUsername_Call struct {
	*mock.Call
}

// HasUsername is a helper method to define mock.On call
//   - username string
func (_e *UserRepository_Expecter) HasUsername(username interface{}) *UserRepository_HasUsername_Call {
	return &UserRepository_HasUsername_Call{Call: _e.mock.On("HasUsername", username)}
}

func (_c *UserRepository_HasUsername_Call) Run(run func(username string)) *UserRepository_HasUsername_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *UserRepository_HasUsername_Call) Return(_a0 bool, _a1 error) *UserRepository_HasUsername_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *UserRepository_HasUsername_Call) RunAndReturn(run func(string) (bool, error)) *UserRepository_HasUsername_Call {
	_c.Call.Return(run)
	return _c
}

// NewUserRepository creates a new instance of UserRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepository {
	mock := &UserRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
