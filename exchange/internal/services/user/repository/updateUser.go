package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	userRepositoryModel "exchange/internal/services/user/model"
	"exchange/internal/services/user/repository/userDDL"
)

// UpdateUser редактирует пользователя
func (r *UserRepository) UpdateUser(ctx context.Context, req userRepositoryModel.UpdateUserReq) error {

	// Редактируем пользователя
	return r.pgsql.Exec(ctx, sq.
		Update(userDDL.Table).
		SetMap(map[string]any{
			userDDL.ColumnLastName:  req.LastName,
			userDDL.ColumnFirstName: req.FirstName,
			userDDL.ColumnEmail:     req.Email,
			userDDL.ColumnAuthorID:  req.AuthorID,
		}).
		Where(sq.Eq{userDDL.ColumnID: req.ID}),
	)
}
