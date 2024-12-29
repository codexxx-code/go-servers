package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/user/model"
	"exchange/internal/services/user/repository/userDDL"
)

// DeleteUser удаляет пользователя
func (r *UserRepository) DeleteUser(ctx context.Context, req model.DeleteUserReq) error {

	// Меняем флаг удаления пользователя на true
	return r.pgsql.Exec(ctx, sq.
		Update(userDDL.Table).
		Set(userDDL.ColumnIsDeleted, true).
		Where(sq.Eq{userDDL.ColumnID: req.ID}),
	)
}
