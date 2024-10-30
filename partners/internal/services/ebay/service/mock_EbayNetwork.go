// Code generated by mockery v2.46.2. DO NOT EDIT.

package service

import (
	context "context"
	model "partners/internal/services/ebay/network/model"

	mock "github.com/stretchr/testify/mock"
)

// MockEbayNetwork is an autogenerated mock type for the EbayNetwork type
type MockEbayNetwork struct {
	mock.Mock
}

// GetCategories provides a mock function with given fields: ctx, req
func (_m *MockEbayNetwork) GetCategories(ctx context.Context, req model.GetCategoriesReq) (model.GetCategoriesRes, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetCategories")
	}

	var r0 model.GetCategoriesRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetCategoriesReq) (model.GetCategoriesRes, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetCategoriesReq) model.GetCategoriesRes); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.GetCategoriesRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetCategoriesReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCategoryTreeID provides a mock function with given fields: ctx, req
func (_m *MockEbayNetwork) GetCategoryTreeID(ctx context.Context, req model.GetCategoryTreeIDReq) (model.GetCategoryTreeIDRes, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetCategoryTreeID")
	}

	var r0 model.GetCategoryTreeIDRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetCategoryTreeIDReq) (model.GetCategoryTreeIDRes, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetCategoryTreeIDReq) model.GetCategoryTreeIDRes); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.GetCategoryTreeIDRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetCategoryTreeIDReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetItems provides a mock function with given fields: ctx, req
func (_m *MockEbayNetwork) GetItems(ctx context.Context, req model.GetItemsReq) (model.GetItemsRes, error) {
	ret := _m.Called(ctx, req)

	if len(ret) == 0 {
		panic("no return value specified for GetItems")
	}

	var r0 model.GetItemsRes
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, model.GetItemsReq) (model.GetItemsRes, error)); ok {
		return rf(ctx, req)
	}
	if rf, ok := ret.Get(0).(func(context.Context, model.GetItemsReq) model.GetItemsRes); ok {
		r0 = rf(ctx, req)
	} else {
		r0 = ret.Get(0).(model.GetItemsRes)
	}

	if rf, ok := ret.Get(1).(func(context.Context, model.GetItemsReq) error); ok {
		r1 = rf(ctx, req)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewMockEbayNetwork creates a new instance of MockEbayNetwork. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockEbayNetwork(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockEbayNetwork {
	mock := &MockEbayNetwork{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
