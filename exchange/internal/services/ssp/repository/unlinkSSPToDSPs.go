package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspToDSPDDL"
)

func (r *SSPRepository) UnlinkSSPToDSPs(ctx context.Context, req model.UnlinkSSPToDSPsReq) error {

	// Выполнение SQL запроса для софт удаления связки между SSP и DSP
	return r.pgsql.Exec(ctx, sq.
		Update(sspToDSPDDL.Table).
		SetMap(map[string]any{
			sspToDSPDDL.ColumnIsDeleted: true,
		}).
		Where(sq.Eq{
			sspToDSPDDL.ColumnSSPSlug: req.SSPSlug,
			sspToDSPDDL.ColumnDSPSlug: req.DSPSlugs,
		}),
	)
}
