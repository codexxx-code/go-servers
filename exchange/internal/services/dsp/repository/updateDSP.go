package repository

import (
	"context"

	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/dsp/model"
	"exchange/internal/services/dsp/repository/dspDDL"
)

func (r *DSPRepository) UpdateDSP(ctx context.Context, req model.UpdateDSPReq) error {

	// Выполнение SQL запроса для обновления данных
	return r.pgsql.Exec(ctx, sq.
		Update(dspDDL.Table).
		SetMap(map[string]any{
			dspDDL.ColumnName:                     req.Name,
			dspDDL.ColumnUrl:                      req.EndpointURL,
			dspDDL.ColumnCurrency:                 req.Currency,
			dspDDL.ColumnAuctionSecondPrice:       req.AuctionSecondPrice,
			dspDDL.ColumnIsSupportMultiimpression: req.IsSupportMultiimpression,
			dspDDL.ColumnBillingUrlType:           req.BillingType,
			dspDDL.ColumnIsEnable:                 req.IsEnable,
			dspDDL.ColumnIntegrationType:          req.IntegrationType,
			dspDDL.ColumnSourceTrafficTypes:       pq.Array(req.SourceTrafficTypes),
			dspDDL.ColumnFormatTypes:              pq.Array(req.FormatTypes),
		}).
		Where(sq.Eq{dspDDL.ColumnSlug: req.Slug}),
	)
}
