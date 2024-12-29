package service

import (
	"context"
	"time"

	"exchange/internal/config"
	"exchange/internal/services/auth/model"
	userModel "exchange/internal/services/user/model"
	"exchange/internal/services/user/model/userFilters"
	"pkg/errors"
	"pkg/passwordManager"
	"pkg/slices"
)

var ErrInvalidLoginOrPassword = errors.New("Invalid login or password.")

// SignIn авторизует пользователя и возвращает токены доступа
func (s *AuthService) SignIn(ctx context.Context, req model.SignInReq) (res model.SignInRes, err error) {

	// Получаем пользователя из базы данных
	findUsersRes, err := s.userService.FindUsers(ctx,
		userModel.FindUsersReq{ //nolint:exhaustruct
			Filters: userFilters.UserFilters{ //nolint:exhaustruct
				Emails: []string{req.Login},
			},
		},
	)
	user, err := slices.FirstWithError(findUsersRes.Users, err)
	if err != nil {
		return res, err
	}

	// Сравниваем пароли
	err = passwordManager.CompareHashAndPassword(
		user.PasswordHash,
		[]byte(req.Password),
		user.PasswordSalt,
		[]byte(config.Load().Auth.GeneralSalt),
	)
	if err != nil {
		return res, errors.BadRequest.Wrap(ErrInvalidLoginOrPassword)
	}

	// Создаем пару токенов доступа
	tokens, err := createPairTokens(user.Permissions, user.ID)
	if err != nil {
		return res, err
	}

	// Обновляем время последнего логина
	err = s.userService.UpdateLastLoginAt(
		ctx,
		userModel.UpdateLastLoginAtReq{
			ID:          user.ID,
			LastLoginAt: time.Time{}, // Заполняется в репозитории для простоты тестирования
		},
	)
	if err != nil {
		return res, err
	}

	return model.SignInRes{
		Tokens: tokens,
		ID:     user.ID,
	}, nil
}
