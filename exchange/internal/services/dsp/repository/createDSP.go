package repository

import (
	"context"

	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"

	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/dsp/repository/dspDDL"
)

func (r *DSPRepository) CreateDSP(ctx context.Context, req dspModel.CreateDSPReq) error {

	// Выполнение SQL запроса для вставки данных
	return r.pgsql.Exec(ctx, sq.
		Insert(dspDDL.Table).
		SetMap(map[string]any{
			dspDDL.ColumnSlug:                     req.Slug,
			dspDDL.ColumnName:                     req.Name,
			dspDDL.ColumnIsSupportMultiimpression: req.IsSupportMultiimpression,
			dspDDL.ColumnUrl:                      req.EndpointURL,
			dspDDL.ColumnCurrency:                 req.Currency,
			dspDDL.ColumnAuctionSecondPrice:       req.AuctionSecondPrice,
			dspDDL.ColumnBillingUrlType:           req.BillingType,
			dspDDL.ColumnIsEnable:                 req.IsEnable,
			dspDDL.ColumnIntegrationType:          req.IntegrationType,
			dspDDL.ColumnSourceTrafficTypes:       pq.Array(req.SourceTrafficTypes),
			dspDDL.ColumnFormatTypes:              pq.Array(req.FormatTypes),
		}),
	)
}
