// Code generated by mockery v2.38.0. DO NOT EDIT.

package hello

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Greeter is an autogenerated mock type for the Greeter type
type Greeter struct {
	mock.Mock
}

type Greeter_Expecter struct {
	mock *mock.Mock
}

func (_m *Greeter) EXPECT() *Greeter_Expecter {
	return &Greeter_Expecter{mock: &_m.Mock}
}

// Execute provides a mock function with given fields: _a0, _a1
func (_m *Greeter) Execute(_a0 context.Context, _a1 string) (string, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Greeter_Execute_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Execute'
type Greeter_Execute_Call struct {
	*mock.Call
}

// Execute is a helper method to define mock.On call
//   - _a0 context.Context
//   - _a1 string
func (_e *Greeter_Expecter) Execute(_a0 interface{}, _a1 interface{}) *Greeter_Execute_Call {
	return &Greeter_Execute_Call{Call: _e.mock.On("Execute", _a0, _a1)}
}

func (_c *Greeter_Execute_Call) Run(run func(_a0 context.Context, _a1 string)) *Greeter_Execute_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *Greeter_Execute_Call) Return(_a0 string, _a1 error) *Greeter_Execute_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *Greeter_Execute_Call) RunAndReturn(run func(context.Context, string) (string, error)) *Greeter_Execute_Call {
	_c.Call.Return(run)
	return _c
}

// NewGreeter creates a new instance of Greeter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewGreeter(t interface {
	mock.TestingT
	Cleanup(func())
}) *Greeter {
	mock := &Greeter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
