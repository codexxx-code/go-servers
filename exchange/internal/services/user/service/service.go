package service

import (
	"context"

	userModel "exchange/internal/services/user/model"
	userRepository "exchange/internal/services/user/repository"
)

var _ UserRepository = new(userRepository.UserRepository)

type UserRepository interface {
	CreateUser(context.Context, userRepository.CreateUserReq) (string, error)
	DeleteUser(context.Context, userModel.DeleteUserReq) error
	GetUsersCount(context.Context, userModel.FindUsersReq) (int, error)
	UpdateLastLoginAt(context.Context, userModel.UpdateLastLoginAtReq) error
	FindUsers(context.Context, userModel.FindUsersReq) ([]userModel.User, error)
	UpdateUserPermissions(context.Context, userModel.UpdateUserPermissionsReq) error
	UpdateUser(context.Context, userModel.UpdateUserReq) error
}

type UserService struct {
	userRepository UserRepository
}

func NewUserService(
	userRepository UserRepository,
	ignoreCreatingRootAdmin bool,
) (*UserService, error) {
	s := &UserService{
		userRepository: userRepository,
	}

	// Инициализируем главного пользователя
	if !ignoreCreatingRootAdmin {
		_, err := s.initializeRootAdmin(context.Background())
		if err != nil {
			return s, err
		}
	}

	return s, nil
}
