package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspDDL"
)

func (r *SSPRepository) DeleteSSP(ctx context.Context, req model.DeleteSSPReq) error {

	// Выполнение SQL запроса для удаления записи по Slug
	return r.pgsql.Exec(ctx, sq.
		Delete(sspDDL.Table).
		Where(sq.Eq{sspDDL.ColumnSlug: req.Slug}),
	)
}
