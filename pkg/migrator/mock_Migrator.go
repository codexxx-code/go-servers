// Code generated by mockery v2.46.2. DO NOT EDIT.

package migrator

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// MockMigrator is an autogenerated mock type for the Migrator type
type MockMigrator struct {
	mock.Mock
}

// Down provides a mock function with given fields: _a0
func (_m *MockMigrator) Down(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Down")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Up provides a mock function with given fields: _a0
func (_m *MockMigrator) Up(_a0 context.Context) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for Up")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewMockMigrator creates a new instance of MockMigrator. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockMigrator(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockMigrator {
	mock := &MockMigrator{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}