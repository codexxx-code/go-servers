package service

import (
	"context"

	userModel "exchange/internal/services/user/model"
)

func (s *UserService) GetUsersCount(ctx context.Context, req userModel.FindUsersReq) (int, error) {
	return s.userRepository.GetUsersCount(ctx, req)
}
