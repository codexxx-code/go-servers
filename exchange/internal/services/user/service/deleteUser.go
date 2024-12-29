package service

import (
	"context"

	"exchange/internal/services/user/model"
)

// DeleteUser удаляет пользователя
func (s *UserService) DeleteUser(ctx context.Context, req model.DeleteUserReq) error {
	return s.userRepository.DeleteUser(ctx, req)
}
