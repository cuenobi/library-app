// Code generated by mockery v2.52.2. DO NOT EDIT.

package mocks

import (
	model "library-service/internal/model"

	mock "github.com/stretchr/testify/mock"
)

// BookRepository is an autogenerated mock type for the BookRepository type
type BookRepository struct {
	mock.Mock
}

type BookRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *BookRepository) EXPECT() *BookRepository_Expecter {
	return &BookRepository_Expecter{mock: &_m.Mock}
}

// CreateBook provides a mock function with given fields: book
func (_m *BookRepository) CreateBook(book *model.Book) error {
	ret := _m.Called(book)

	if len(ret) == 0 {
		panic("no return value specified for CreateBook")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*model.Book) error); ok {
		r0 = rf(book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BookRepository_CreateBook_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateBook'
type BookRepository_CreateBook_Call struct {
	*mock.Call
}

// CreateBook is a helper method to define mock.On call
//   - book *model.Book
func (_e *BookRepository_Expecter) CreateBook(book interface{}) *BookRepository_CreateBook_Call {
	return &BookRepository_CreateBook_Call{Call: _e.mock.On("CreateBook", book)}
}

func (_c *BookRepository_CreateBook_Call) Run(run func(book *model.Book)) *BookRepository_CreateBook_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*model.Book))
	})
	return _c
}

func (_c *BookRepository_CreateBook_Call) Return(_a0 error) *BookRepository_CreateBook_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BookRepository_CreateBook_Call) RunAndReturn(run func(*model.Book) error) *BookRepository_CreateBook_Call {
	_c.Call.Return(run)
	return _c
}

// DecreaseBookStock provides a mock function with given fields: id
func (_m *BookRepository) DecreaseBookStock(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DecreaseBookStock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BookRepository_DecreaseBookStock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DecreaseBookStock'
type BookRepository_DecreaseBookStock_Call struct {
	*mock.Call
}

// DecreaseBookStock is a helper method to define mock.On call
//   - id string
func (_e *BookRepository_Expecter) DecreaseBookStock(id interface{}) *BookRepository_DecreaseBookStock_Call {
	return &BookRepository_DecreaseBookStock_Call{Call: _e.mock.On("DecreaseBookStock", id)}
}

func (_c *BookRepository_DecreaseBookStock_Call) Run(run func(id string)) *BookRepository_DecreaseBookStock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BookRepository_DecreaseBookStock_Call) Return(_a0 error) *BookRepository_DecreaseBookStock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BookRepository_DecreaseBookStock_Call) RunAndReturn(run func(string) error) *BookRepository_DecreaseBookStock_Call {
	_c.Call.Return(run)
	return _c
}

// GetBookByID provides a mock function with given fields: id
func (_m *BookRepository) GetBookByID(id string) (*model.Book, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetBookByID")
	}

	var r0 *model.Book
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*model.Book, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *model.Book); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Book)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BookRepository_GetBookByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetBookByID'
type BookRepository_GetBookByID_Call struct {
	*mock.Call
}

// GetBookByID is a helper method to define mock.On call
//   - id string
func (_e *BookRepository_Expecter) GetBookByID(id interface{}) *BookRepository_GetBookByID_Call {
	return &BookRepository_GetBookByID_Call{Call: _e.mock.On("GetBookByID", id)}
}

func (_c *BookRepository_GetBookByID_Call) Run(run func(id string)) *BookRepository_GetBookByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BookRepository_GetBookByID_Call) Return(_a0 *model.Book, _a1 error) *BookRepository_GetBookByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BookRepository_GetBookByID_Call) RunAndReturn(run func(string) (*model.Book, error)) *BookRepository_GetBookByID_Call {
	_c.Call.Return(run)
	return _c
}

// HasBookName provides a mock function with given fields: name
func (_m *BookRepository) HasBookName(name string) (bool, error) {
	ret := _m.Called(name)

	if len(ret) == 0 {
		panic("no return value specified for HasBookName")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(name)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// BookRepository_HasBookName_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'HasBookName'
type BookRepository_HasBookName_Call struct {
	*mock.Call
}

// HasBookName is a helper method to define mock.On call
//   - name string
func (_e *BookRepository_Expecter) HasBookName(name interface{}) *BookRepository_HasBookName_Call {
	return &BookRepository_HasBookName_Call{Call: _e.mock.On("HasBookName", name)}
}

func (_c *BookRepository_HasBookName_Call) Run(run func(name string)) *BookRepository_HasBookName_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BookRepository_HasBookName_Call) Return(_a0 bool, _a1 error) *BookRepository_HasBookName_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *BookRepository_HasBookName_Call) RunAndReturn(run func(string) (bool, error)) *BookRepository_HasBookName_Call {
	_c.Call.Return(run)
	return _c
}

// IncreaseBookStock provides a mock function with given fields: id
func (_m *BookRepository) IncreaseBookStock(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for IncreaseBookStock")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// BookRepository_IncreaseBookStock_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'IncreaseBookStock'
type BookRepository_IncreaseBookStock_Call struct {
	*mock.Call
}

// IncreaseBookStock is a helper method to define mock.On call
//   - id string
func (_e *BookRepository_Expecter) IncreaseBookStock(id interface{}) *BookRepository_IncreaseBookStock_Call {
	return &BookRepository_IncreaseBookStock_Call{Call: _e.mock.On("IncreaseBookStock", id)}
}

func (_c *BookRepository_IncreaseBookStock_Call) Run(run func(id string)) *BookRepository_IncreaseBookStock_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(string))
	})
	return _c
}

func (_c *BookRepository_IncreaseBookStock_Call) Return(_a0 error) *BookRepository_IncreaseBookStock_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *BookRepository_IncreaseBookStock_Call) RunAndReturn(run func(string) error) *BookRepository_IncreaseBookStock_Call {
	_c.Call.Return(run)
	return _c
}

// NewBookRepository creates a new instance of BookRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewBookRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *BookRepository {
	mock := &BookRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
