// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	fiber "github.com/gofiber/fiber/v2"
	mock "github.com/stretchr/testify/mock"
)

// JWT is an autogenerated mock type for the JWT type
type JWT struct {
	mock.Mock
}

type JWT_Expecter struct {
	mock *mock.Mock
}

func (_m *JWT) EXPECT() *JWT_Expecter {
	return &JWT_Expecter{mock: &_m.Mock}
}

// Generate provides a mock function with given fields: username, role
func (_m *JWT) Generate(username string, role string) string {
	ret := _m.Called(username, role)

	if len(ret) == 0 {
		panic("no return value specified for Generate")
	}

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(username, role)
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// JWT_Generate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Generate'
type JWT_Generate_Call struct {
	*mock.Call
}

// Generate is a helper method to define mock.On call
//   - username string
//   - role string
func (_e *JWT_Expecter) Generate(username interface{}, role interface{}) *JWT_Generate_Call {
	return &JWT_Generate_Call{Call: _e.mock.On("Generate", username, role)}
}

func (_c *JWT_Generate_Call) Run(run func(username string, role string)) *JWT_Generate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string), args[1].(string))
	})
	return _c
}

func (_c *JWT_Generate_Call) Return(_a0 string) *JWT_Generate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *JWT_Generate_Call) RunAndReturn(run func(string, string) string) *JWT_Generate_Call {
	_c.Call.Return(run)
	return _c
}

// Validate provides a mock function with given fields: c
func (_m *JWT) Validate(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for Validate")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// JWT_Validate_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Validate'
type JWT_Validate_Call struct {
	*mock.Call
}

// Validate is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *JWT_Expecter) Validate(c interface{}) *JWT_Validate_Call {
	return &JWT_Validate_Call{Call: _e.mock.On("Validate", c)}
}

func (_c *JWT_Validate_Call) Run(run func(c *fiber.Ctx)) *JWT_Validate_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *JWT_Validate_Call) Return(_a0 error) *JWT_Validate_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *JWT_Validate_Call) RunAndReturn(run func(*fiber.Ctx) error) *JWT_Validate_Call {
	_c.Call.Return(run)
	return _c
}

// ValidateLibrarian provides a mock function with given fields: c
func (_m *JWT) ValidateLibrarian(c *fiber.Ctx) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for ValidateLibrarian")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*fiber.Ctx) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// JWT_ValidateLibrarian_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ValidateLibrarian'
type JWT_ValidateLibrarian_Call struct {
	*mock.Call
}

// ValidateLibrarian is a helper method to define mock.On call
//   - c *fiber.Ctx
func (_e *JWT_Expecter) ValidateLibrarian(c interface{}) *JWT_ValidateLibrarian_Call {
	return &JWT_ValidateLibrarian_Call{Call: _e.mock.On("ValidateLibrarian", c)}
}

func (_c *JWT_ValidateLibrarian_Call) Run(run func(c *fiber.Ctx)) *JWT_ValidateLibrarian_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*fiber.Ctx))
	})
	return _c
}

func (_c *JWT_ValidateLibrarian_Call) Return(_a0 error) *JWT_ValidateLibrarian_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *JWT_ValidateLibrarian_Call) RunAndReturn(run func(*fiber.Ctx) error) *JWT_ValidateLibrarian_Call {
	_c.Call.Return(run)
	return _c
}

// NewJWT creates a new instance of JWT. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJWT(t interface {
	mock.TestingT
	Cleanup(func())
}) *JWT {
	mock := &JWT{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
