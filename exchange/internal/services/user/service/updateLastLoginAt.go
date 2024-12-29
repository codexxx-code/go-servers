package service

import (
	"context"

	"exchange/internal/services/user/model"
)

// UpdateLastLoginAt обновляет время последнего входа
func (s *UserService) UpdateLastLoginAt(ctx context.Context, req model.UpdateLastLoginAtReq) error {
	return s.userRepository.UpdateLastLoginAt(ctx, req)
}
