package repository

import (
	"context"

	"github.com/lib/pq"

	sq "github.com/Masterminds/squirrel"

	sspModel "exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspDDL"
)

func (r *SSPRepository) UpdateSSP(ctx context.Context, req sspModel.UpdateSSPReq) error {

	// Выполнение SQL запроса для обновления данных
	return r.pgsql.Exec(ctx, sq.
		Update(sspDDL.Table).
		SetMap(map[string]any{
			sspDDL.ColumnName:                req.Name,
			sspDDL.ColumnTimeout:             req.Timeout,
			sspDDL.ColumnIsEnable:            req.IsEnable,
			sspDDL.ColumnIntegrationType:     req.IntegrationType,
			sspDDL.ColumnSourceTrafficTypes:  pq.Array(req.SourceTrafficTypes),
			sspDDL.ColumnBillingType:         req.BillingType,
			sspDDL.ColumnAuctionSecondPrice:  req.AuctionSecondPrice,
			sspDDL.ColumnCurrency:            req.Currency,
			sspDDL.ColumnFraudScore:          req.FraudScore,
			sspDDL.ColumnFormatTypes:         pq.Array(req.FormatTypes),
			sspDDL.ColumnClickunderDrumSize:  req.ClickunderDrumSize,
			sspDDL.ColumnClickunderADMFormat: req.ClickunderADMFormat,
		}).
		Where(sq.Eq{sspDDL.ColumnSlug: req.Slug}),
	)
}
