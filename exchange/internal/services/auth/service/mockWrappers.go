package service

import (
	"testing"

	"github.com/stretchr/testify/mock"

	userModel "exchange/internal/services/user/model"
)

type Mocks struct {
	MockUserService *MockUserService
}

func NewMocks(t *testing.T) *Mocks {
	return &Mocks{
		MockUserService: NewMockUserService(t),
	}
}

func (m *Mocks) MockCreateUser(req userModel.CreateUserReq,
	res string, err error,
) {
	m.MockUserService.On("CreateUser", mock.Anything, req).Return(res, err)
}

func (m *Mocks) MockDeleteUser(req userModel.DeleteUserReq,
	err error,
) {
	m.MockUserService.On("DeleteUser", mock.Anything, req).Return(err)
}

func (m *Mocks) MockGetUsersCount(req userModel.FindUsersReq,
	res int, err error,
) {
	m.MockUserService.On("GetUsersCount", mock.Anything, req).Return(res, err)
}

func (m *Mocks) MockUpdateLastLoginAt(req userModel.UpdateLastLoginAtReq,
	err error,
) {
	m.MockUserService.On("UpdateLastLoginAt", mock.Anything, req).Return(err)
}

func (m *Mocks) MockFindUsers(req userModel.FindUsersReq,
	res userModel.FindUsersRes, err error,
) {
	m.MockUserService.On("FindUsers", mock.Anything, req).Return(res, err)
}

func (m *Mocks) MockUpdateUserPermissions(req userModel.UpdateUserPermissionsReq,
	err error,
) {
	m.MockUserService.On("UpdateUserPermissions", mock.Anything, req).Return(err)
}

func (m *Mocks) MockUpdateUser(req userModel.UpdateUserReq,
	err error,
) {
	m.MockUserService.On("UpdateUser", mock.Anything, req).Return(err)
}
