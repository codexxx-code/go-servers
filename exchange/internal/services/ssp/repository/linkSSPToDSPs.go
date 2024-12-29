package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspToDSPDDL"
)

func (r *SSPRepository) LinkSSPToDSPs(ctx context.Context, req model.LinkSSPToDSPsReq) error {

	q := sq.
		Insert(sspToDSPDDL.Table).
		Columns(
			sspToDSPDDL.ColumnSSPSlug,
			sspToDSPDDL.ColumnDSPSlug,
			sspToDSPDDL.ColumnIsDeleted,
		).
		Suffix("ON CONFLICT DO UPDATE SET is_deleted = false")

	for _, dsp := range req.DSPSlugs {
		q = q.Values(req.SSPSlug, dsp, false)
	}

	// Выполнение SQL запроса для создания связки между SSP и DSP
	return r.pgsql.Exec(ctx, q)
}
