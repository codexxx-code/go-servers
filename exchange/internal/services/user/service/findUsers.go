package service

import (
	"context"

	"exchange/internal/services/user/model"
	"pkg/pointer"
)

// FindUsers получает список пользователей
func (s *UserService) FindUsers(ctx context.Context, req model.FindUsersReq) (res model.FindUsersRes, err error) {

	req.Filters.IsDeleted = pointer.Pointer(false)

	// Получаем пользователей
	users, err := s.userRepository.FindUsers(ctx, req)
	if err != nil {
		return res, err
	}

	// Получаем количество пользователей
	usersCount, err := s.userRepository.GetUsersCount(ctx, req)
	if err != nil {
		return res, err
	}

	return model.FindUsersRes{
		Users:      users,
		UsersCount: usersCount,
	}, nil
}
