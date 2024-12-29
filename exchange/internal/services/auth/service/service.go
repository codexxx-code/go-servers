package service

import (
	"context"

	userModel "exchange/internal/services/user/model"
	userService "exchange/internal/services/user/service"
)

var _ UserService = new(userService.UserService)

type UserService interface {
	UpdateLastLoginAt(context.Context, userModel.UpdateLastLoginAtReq) error
	FindUsers(context.Context, userModel.FindUsersReq) (userModel.FindUsersRes, error)
}

type AuthService struct {
	userService UserService
}

func NewAuthService(
	userService UserService,
) *AuthService {
	return &AuthService{
		userService: userService,
	}
}
