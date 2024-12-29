package repository

import (
	"context"
	"time"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/user/model"
	"exchange/internal/services/user/repository/userDDL"
)

// UpdateLastLoginAt обновляет время последнего входа в приложение
func (r *UserRepository) UpdateLastLoginAt(ctx context.Context, req model.UpdateLastLoginAtReq) error {

	// Определяем текущее время
	req.LastLoginAt = time.Now()

	// Исполняем запрос на обновление последнего входа в приложение
	return r.pgsql.Exec(ctx, sq.
		Update(userDDL.Table).
		Set(userDDL.ColumnLastLoginAt, req.LastLoginAt).
		Where(sq.Eq{userDDL.ColumnID: req.ID}),
	)
}
