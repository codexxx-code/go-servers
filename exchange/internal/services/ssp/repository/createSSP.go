package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	sspModel "exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/repository/sspDDL"
)

func (r *SSPRepository) CreateSSP(ctx context.Context, req sspModel.CreateSSPReq) error {

	// Выполнение SQL запроса для вставки данных
	return r.pgsql.Exec(ctx, sq.
		Insert(sspDDL.Table).
		SetMap(map[string]any{
			sspDDL.ColumnSlug:                req.Slug,
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
		}),
	)
}
