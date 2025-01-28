// Code generated by mockery v2.39.1. DO NOT EDIT.

package free_addresses

import mock "github.com/stretchr/testify/mock"

// MockExecuter is an autogenerated mock type for the Executer type
type MockExecuter struct {
	mock.Mock
}

// Execute provides a mock function with given fields: command, args
func (_m *MockExecuter) Execute(command string, args ...string) (string, string, int) {
	_va := make([]interface{}, len(args))
	for _i := range args {
		_va[_i] = args[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, command)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for Execute")
	}

	var r0 string
	var r1 string
	var r2 int
	if rf, ok := ret.Get(0).(func(string, ...string) (string, string, int)); ok {
		return rf(command, args...)
	}
	if rf, ok := ret.Get(0).(func(string, ...string) string); ok {
		r0 = rf(command, args...)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(string, ...string) string); ok {
		r1 = rf(command, args...)
	} else {
		r1 = ret.Get(1).(string)
	}

	if rf, ok := ret.Get(2).(func(string, ...string) int); ok {
		r2 = rf(command, args...)
	} else {
		r2 = ret.Get(2).(int)
	}

	return r0, r1, r2
}

// NewMockExecuter creates a new instance of MockExecuter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockExecuter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockExecuter {
	mock := &MockExecuter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
