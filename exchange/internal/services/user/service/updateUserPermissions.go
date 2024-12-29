package service

import (
	"context"

	"exchange/internal/services/user/model"
)

// UpdateUserPermissions обновляет разрешения
func (s *UserService) UpdateUserPermissions(ctx context.Context, req model.UpdateUserPermissionsReq) (err error) {
	return s.userRepository.UpdateUserPermissions(ctx, req)
}
