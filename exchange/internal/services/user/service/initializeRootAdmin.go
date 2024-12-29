package service

import (
	"context"
	"fmt"

	"exchange/internal/config"
	"exchange/internal/services/user/model"
	"exchange/internal/services/user/model/userFilters"
	userRepository "exchange/internal/services/user/repository"
	"pkg/log"

	"pkg/slices"

	"exchange/internal/enum/permission"
	"pkg/errors"
	"pkg/passwordManager"
)

func (s *UserService) initializeRootAdmin(ctx context.Context) (id string, err error) {

	rootAdminLogin := config.Load().Auth.RootAdminLogin

	// Ищем пользователя по электронной почте
	user, err := slices.FirstWithError(s.userRepository.FindUsers(ctx,
		model.FindUsersReq{ //nolint:exhaustruct
			Filters: userFilters.UserFilters{ //nolint:exhaustruct
				Emails: []string{rootAdminLogin},
			},
		},
	))
	if err != nil && !errors.Is(err, slices.ErrSliceIsEmpty) {
		return id, err
	}

	// Если ошибка пустая, значит пользователь найден
	if err == nil {
		log.Info(ctx, fmt.Sprintf("Root admin already exists. UserID: %s", user.ID))
		return user.ID, nil
	}

	// Генерируем пароль
	passwordSalt, err := passwordManager.GenerateRandomSalt()
	if err != nil {
		return id, err
	}
	passwordHash, err := passwordManager.CreateNewPassword(
		[]byte(config.Load().Auth.RootAdminPassword),
		[]byte(config.Load().Auth.GeneralSalt),
		passwordSalt,
	)
	if err != nil {
		return id, err
	}

	id, err = s.userRepository.CreateUser(ctx,
		userRepository.CreateUserReq{
			ID:           "",
			LastName:     "",
			FirstName:    "",
			Email:        rootAdminLogin,
			PasswordHash: passwordHash,
			PasswordSalt: passwordSalt,
			Permissions: []permission.Permission{
				permission.Root,
			},
			AuthorID: nil,
		},
	)
	if err != nil {
		return id, err
	}

	log.Info(ctx, fmt.Sprintf("Root admin created. UserID: %s", id))

	return id, nil
}
