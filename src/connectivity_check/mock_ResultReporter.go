// Code generated by mockery v2.39.1. DO NOT EDIT.

package connectivity_check

import (
	models "github.com/openshift/assisted-service/models"
	mock "github.com/stretchr/testify/mock"
)

// MockResultReporter is an autogenerated mock type for the ResultReporter type
type MockResultReporter struct {
	mock.Mock
}

// Report provides a mock function with given fields: resultingHost
func (_m *MockResultReporter) Report(resultingHost *models.ConnectivityRemoteHost) error {
	ret := _m.Called(resultingHost)

	if len(ret) == 0 {
		panic("no return value specified for Report")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.ConnectivityRemoteHost) error); ok {
		r0 = rf(resultingHost)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockResultReporter creates a new instance of MockResultReporter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockResultReporter(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockResultReporter {
	mock := &MockResultReporter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
