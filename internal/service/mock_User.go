// Code generated by mockery v2.14.0. DO NOT EDIT.

package service

import (
	context "context"

	model "github.com/albertopformoso/inventory/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// MockUser is an autogenerated mock type for the User type
type MockUser struct {
	mock.Mock
}

// LoginUser provides a mock function with given fields: ctx, email, password
func (_m *MockUser) LoginUser(ctx context.Context, email string, password string) (*model.User, error) {
	ret := _m.Called(ctx, email, password)

	var r0 *model.User
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *model.User); ok {
		r0 = rf(ctx, email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, string) error); ok {
		r1 = rf(ctx, email, password)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RegisterUser provides a mock function with given fields: ctx, email, name, password
func (_m *MockUser) RegisterUser(ctx context.Context, email string, name string, password string) error {
	ret := _m.Called(ctx, email, name, password)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, email, name, password)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewMockUser interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockUser creates a new instance of MockUser. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockUser(t mockConstructorTestingTNewMockUser) *MockUser {
	mock := &MockUser{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
