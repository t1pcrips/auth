// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	model "github.com/t1pcrips/auth/internal/model"
)

// UserService is an autogenerated mock type for the UserService type
type UserService struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, info
func (_m *UserService) Create(ctx context.Context, info *model.CreateUserRequest) (int64, error) {
	ret := _m.Called(ctx, info)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateUserRequest) (int64, error)); ok {
		return rf(ctx, info)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CreateUserRequest) int64); ok {
		r0 = rf(ctx, info)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CreateUserRequest) error); ok {
		r1 = rf(ctx, info)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, userId
func (_m *UserService) Delete(ctx context.Context, userId int64) error {
	ret := _m.Called(ctx, userId)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, userId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: ctx, chatId
func (_m *UserService) Get(ctx context.Context, chatId int64) (*model.GetUserResponse, error) {
	ret := _m.Called(ctx, chatId)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *model.GetUserResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*model.GetUserResponse, error)); ok {
		return rf(ctx, chatId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *model.GetUserResponse); ok {
		r0 = rf(ctx, chatId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.GetUserResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, chatId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, info
func (_m *UserService) Update(ctx context.Context, info *model.UpdatUsereRequest) error {
	ret := _m.Called(ctx, info)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UpdatUsereRequest) error); ok {
		r0 = rf(ctx, info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserService creates a new instance of UserService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserService(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserService {
	mock := &UserService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
