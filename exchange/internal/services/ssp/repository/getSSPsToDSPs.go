package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspToDSPDDL"
	"pkg/ddlHelper"
)

func (r *SSPRepository) getSSPsToDSPs(ctx context.Context, sspSlugs []string) (dsps []model.SSPToDSP, err error) {

	filters := sq.Eq{
		sspToDSPDDL.ColumnIsDeleted: false,
	}

	if len(sspSlugs) != 0 {
		filters[sspToDSPDDL.ColumnSSPSlug] = sspSlugs
	}

	// Выполнение SQL запроса для получения связок между SSP и DSP
	return dsps, r.pgsql.Select(ctx, &dsps, sq.
		Select(ddlHelper.SelectAll).
		From(sspToDSPDDL.Table).
		Where(filters),
	)
}
