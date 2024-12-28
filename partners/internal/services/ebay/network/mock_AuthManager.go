// Code generated by mockery v2.46.2. DO NOT EDIT.

package network

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockAuthManager is an autogenerated mock type for the AuthManager type
type MockAuthManager struct {
	mock.Mock
}

// GetToken provides a mock function with given fields: ctx
func (_m *MockAuthManager) GetToken(ctx context.Context) (string, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for GetToken")
	}

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (string, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockAuthManager creates a new instance of MockAuthManager. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAuthManager(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAuthManager {
	mock := &MockAuthManager{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}