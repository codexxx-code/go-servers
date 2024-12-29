package service

import (
	"context"

	"exchange/internal/services/user/model"
)

// UpdateUser обновляет пользователя
func (s *UserService) UpdateUser(ctx context.Context, req model.UpdateUserReq) error {
	return s.userRepository.UpdateUser(ctx, req)
}
