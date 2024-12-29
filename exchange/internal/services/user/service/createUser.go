package service

import (
	"context"
	"strings"

	"exchange/internal/config"
	userModel "exchange/internal/services/user/model"
	"exchange/internal/services/user/model/userFilters"
	userRepository "exchange/internal/services/user/repository"
	"pkg/passwordManager"
	"pkg/slices"
	"pkg/uuid"

	"pkg/errors"
)

var ErrPasswordsNotMatch = errors.New("Passwords don't match")
var ErrUserAlreadyExist = errors.New("The user already exists")

// CreateUser создает пользователя
func (s *UserService) CreateUser(ctx context.Context, req userModel.CreateUserReq) (id string, err error) {

	// Проверяем, чтобы пароль и подтверждение пароля были одинаковыми
	if req.Password != req.RetryPassword {
		return id, errors.BadRequest.Wrap(ErrPasswordsNotMatch)
	}

	// Переводим значения электронной почты в строчные буквы
	req.Email = strings.ToLower(req.Email)

	// Пытаемся найти пользователя с такой почтой
	_, err = slices.FirstWithError(
		s.userRepository.FindUsers(
			ctx, userModel.FindUsersReq{ //nolint:exhaustruct
				Filters: userFilters.UserFilters{ //nolint:exhaustruct
					Emails: []string{req.Email},
				},
			},
		),
	)
	if err != nil && !errors.Is(err, slices.ErrSliceIsEmpty) {
		return id, err
	}

	// Если ошибка пустая, значит нашли пользователя в базе данных
	if err == nil {
		return id, errors.Forbidden.Wrap(ErrUserAlreadyExist)
	}

	// Хэшируем пароль
	passwordSalt, err := passwordManager.GenerateRandomSalt()
	if err != nil {
		return id, err
	}
	passwordHash, err := passwordManager.CreateNewPassword(
		[]byte(req.Password),
		[]byte(config.Load().Auth.GeneralSalt),
		passwordSalt,
	)
	if err != nil {
		return id, err
	}

	// Создаем пользователя
	return s.userRepository.CreateUser(ctx, userRepository.CreateUserReq{
		ID:           uuid.New(),
		LastName:     req.LastName,
		FirstName:    req.FirstName,
		Email:        req.Email,
		PasswordHash: passwordHash,
		PasswordSalt: passwordSalt,
		Permissions:  req.Permissions,
		AuthorID:     req.AuthorID,
	})
}
