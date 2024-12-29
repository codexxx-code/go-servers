package service

import (
	"testing"

	"github.com/stretchr/testify/mock"

	userModel "exchange/internal/services/user/model"
)

type Mocks struct {
	UserRepository *MockUserRepository
}

func NewMocks(t *testing.T) *Mocks {
	return &Mocks{
		UserRepository: NewMockUserRepository(t),
	}
}

func (m *MockUserRepository) MockCreateUser( // TODO: Сделать Hasher интерфейсом и передавать сюда заранее смоканные данные
	res string, err error,
) {
	m.On("CreateUser", mock.Anything, mock.Anything).Return(res, err)
}

func (m *MockUserRepository) MockDeleteUser(req userModel.DeleteUserReq,
	err error,
) {
	m.On("DeleteUser", mock.Anything, req).Return(err)
}

func (m *MockUserRepository) MockGetUsersCount(req userModel.FindUsersReq,
	res int, err error,
) {
	m.On("GetUsersCount", mock.Anything, req).Return(res, err)
}

func (m *MockUserRepository) MockUpdateLastLoginAt(req userModel.UpdateLastLoginAtReq,
	err error,
) {
	m.On("UpdateLastLoginAt", mock.Anything, req).Return(err)
}

func (m *MockUserRepository) MockFindUsers(req userModel.FindUsersReq,
	res []userModel.User, err error,
) {
	m.On("FindUsers", mock.Anything, req).Return(res, err)
}

func (m *MockUserRepository) MockFindUser(req userModel.FindUsersReq,
	res userModel.User, err error,
) {
	m.On("FindUsers", mock.Anything, req).Return([]userModel.User{res}, err)
}

func (m *MockUserRepository) MockUpdateUserPermissions(req userModel.UpdateUserPermissionsReq,
	err error,
) {
	m.On("UpdateUserPermissions", mock.Anything, req).Return(err)
}

func (m *MockUserRepository) MockUpdateUser(req userModel.UpdateUserReq,
	err error,
) {
	m.On("UpdateUser", mock.Anything, req).Return(err)
}
