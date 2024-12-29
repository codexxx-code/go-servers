package service

import (
	"context"

	"exchange/internal/services/auth/model"
	userModel "exchange/internal/services/user/model"
	"exchange/internal/services/user/model/userFilters"
	"pkg/jwtManager"
	"pkg/slices"
)

// RefreshTokens обновляет токены доступа в базе данных
func (s *AuthService) RefreshTokens(ctx context.Context, request model.RefreshTokensReq) (tokens model.Tokens, err error) {

	// Парсим токен
	claims, err := jwtManager.ParseToken[model.Claims](request.RefreshToken)
	if err != nil {
		return tokens, err
	}

	// Ищем пользователя по идентификатору
	res, err := s.userService.FindUsers(
		ctx,
		userModel.FindUsersReq{ //nolint:exhaustruct
			Filters: userFilters.UserFilters{ //nolint:exhaustruct
				IDs: []string{claims.ID},
			},
		},
	)
	user, err := slices.FirstWithError(res.Users, err)
	if err != nil {
		return tokens, err
	}

	// Создаем новую пару токенов
	tokens, err = createPairTokens(user.Permissions, user.ID)
	if err != nil {
		return tokens, err
	}

	// Возвращаем пару токенов клиенту
	return tokens, nil
}
