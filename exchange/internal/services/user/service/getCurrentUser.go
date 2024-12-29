package service

import (
	"context"

	userModel "exchange/internal/services/user/model"
	"exchange/internal/services/user/model/userFilters"
	"pkg/slices"
)

// GetCurrentUser ищет пользователя по идентификатору
func (s *UserService) GetCurrentUser(ctx context.Context, id string) (userModel.User, error) {
	return slices.FirstWithError(
		s.userRepository.FindUsers(ctx,
			userModel.FindUsersReq{ //nolint:exhaustruct
				Filters: userFilters.UserFilters{
					IDs: []string{id},
				},
			},
		),
	)
}
