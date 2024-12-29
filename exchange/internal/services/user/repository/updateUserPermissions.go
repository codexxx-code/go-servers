package repository

import (
	"context"
	"exchange/internal/services/user/repository/userDDL"
	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/user/model"
)

// UpdateUserPermissions редактирует пользователя
func (r *UserRepository) UpdateUserPermissions(ctx context.Context, req model.UpdateUserPermissionsReq) error {

	// Выполнение SQL запроса для обновления данных
	return r.pgsql.Exec(ctx, sq.
		Update(userDDL.Table).
		SetMap(map[string]any{
			userDDL.ColumnPermissions: pq.Array(req.Permissions),
		}).
		Where(sq.Eq{userDDL.ColumnID: req.ID}),
	)
}
