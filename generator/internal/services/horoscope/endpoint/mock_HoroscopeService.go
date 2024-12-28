// Code generated by mockery v2.46.2. DO NOT EDIT.

package endpoint

import (
	context "context"
	model "generator/internal/services/horoscope/model"

	mock "github.com/stretchr/testify/mock"
)

// MockHoroscopeService is an autogenerated mock type for the HoroscopeService type
type MockHoroscopeService struct {
	mock.Mock
}

// GetHoroscope provides a mock function with given fields: _a0, _a1
func (_m *MockHoroscopeService) GetHoroscope(_a0 context.Context, _a1 model.GetHoroscopeReq) (model.Horoscope, error) {
	ret := _m.Called(_a0, _a1)

	if len(ret) == 0 {
		panic("no return value specified for GetHoroscope")
	}

	var r0 model.Horoscope
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetHoroscopeReq) (model.Horoscope, error)); ok {
		return rf(_a0, _a1)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetHoroscopeReq) model.Horoscope); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(model.Horoscope)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetHoroscopeReq) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockHoroscopeService creates a new instance of MockHoroscopeService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockHoroscopeService(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockHoroscopeService {
	mock := &MockHoroscopeService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}