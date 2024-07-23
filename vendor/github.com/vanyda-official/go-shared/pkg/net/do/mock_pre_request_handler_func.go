// Code generated by mockery. DO NOT EDIT.

package do

import (
	context "context"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockPreRequestHandlerFunc is an autogenerated mock type for the PreRequestHandlerFunc type
type MockPreRequestHandlerFunc struct {
	mock.Mock
}

type MockPreRequestHandlerFunc_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPreRequestHandlerFunc) EXPECT() *MockPreRequestHandlerFunc_Expecter {
	return &MockPreRequestHandlerFunc_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: ctx, request
func (_m *MockPreRequestHandlerFunc) Execute(ctx context.Context, request *http.Request) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *http.Request) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPreRequestHandlerFunc_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type MockPreRequestHandlerFunc_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - ctx context.Context
//   - request *http.Request
func (_e *MockPreRequestHandlerFunc_Expecter) Execute(ctx interface{}, request interface{}) *MockPreRequestHandlerFunc_Execute_Call {
	return &MockPreRequestHandlerFunc_Execute_Call{Call: _e.mock.On("Execute", ctx, request)}
}

func (_c *MockPreRequestHandlerFunc_Execute_Call) Run(run func(ctx context.Context, request *http.Request)) *MockPreRequestHandlerFunc_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*http.Request))
	})
	return _c
}

func (_c *MockPreRequestHandlerFunc_Execute_Call) Return(_a0 error) *MockPreRequestHandlerFunc_Execute_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPreRequestHandlerFunc_Execute_Call) RunAndReturn(run func(context.Context, *http.Request) error) *MockPreRequestHandlerFunc_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPreRequestHandlerFunc creates a new instance of MockPreRequestHandlerFunc. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPreRequestHandlerFunc(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPreRequestHandlerFunc {
	mock := &MockPreRequestHandlerFunc{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}