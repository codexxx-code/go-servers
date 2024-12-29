package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/dsp/model"
	"exchange/internal/services/dsp/repository/dspDDL"
)

func (r *DSPRepository) DeleteDSP(ctx context.Context, req model.DeleteDSPReq) error {

	// Выполнение SQL запроса для удаления записи по Slug
	return r.pgsql.Exec(ctx, sq.
		Delete(dspDDL.Table).
		Where(sq.Eq{dspDDL.ColumnSlug: req.Slug}),
	)
}
