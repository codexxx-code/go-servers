package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"exchange/internal/services/ssp/model/sspFilters"
	"exchange/internal/services/ssp/repository/sspDDL"
	"pkg/ddlHelper"
)

func GetSSPsFilters(filters sspFilters.SSPFilters, q squirrel.SelectBuilder) squirrel.SelectBuilder {
	// Фильтрация по Slugs
	if len(filters.Slugs) > 0 {
		q = q.Where(squirrel.Eq{sspDDL.ColumnSlug: filters.Slugs})
	}

	// Фильтрация по Names
	if len(filters.Names) > 0 {
		q = q.Where(squirrel.Eq{sspDDL.ColumnName: filters.Names})
	}

	// Фильтрация по IsEnable
	if filters.IsEnable != nil {
		q = q.Where(squirrel.Eq{sspDDL.ColumnIsEnable: *filters.IsEnable})
	}

	// Фильтрация по IntegrationTypes
	if len(filters.IntegrationTypes) > 0 {
		q = q.Where(squirrel.Eq{sspDDL.ColumnIntegrationType: filters.IntegrationTypes})
	}

	// Фильтрация по SourceTrafficTypes
	if len(filters.SourceTrafficTypes) > 0 {
		q = q.Where(ddlHelper.PartContains(sspDDL.ColumnSourceTrafficTypes, pq.Array(filters.SourceTrafficTypes)))
	}

	// Фильтрация по BillingTypes
	if len(filters.BillingTypes) > 0 {
		q = q.Where(squirrel.Eq{sspDDL.ColumnBillingType: filters.BillingTypes})
	}

	// Фильтрация по AuctionSecondPrice
	if filters.AuctionSecondPrice != nil {
		q = q.Where(squirrel.Eq{sspDDL.ColumnAuctionSecondPrice: *filters.AuctionSecondPrice})
	}

	// Фильтрация по Currencies
	if len(filters.Currencies) > 0 {
		q = q.Where(squirrel.Eq{sspDDL.ColumnCurrency: filters.Currencies})
	}

	// Фильтрация по FormatTypes
	if len(filters.FormatTypes) > 0 {
		q = q.Where(ddlHelper.PartContains(sspDDL.ColumnFormatTypes, pq.Array(filters.FormatTypes)))
	}

	// Фильтрация по ClickunderDrumSize
	if filters.ClickunderDrumSize != nil {
		q = q.Where(squirrel.Eq{sspDDL.ColumnClickunderDrumSize: *filters.ClickunderDrumSize})
	}

	// Фильтрация по ClickunderADMFormat
	if filters.ClickunderADMFormat != nil {
		q = q.Where(squirrel.Eq{sspDDL.ColumnClickunderADMFormat: *filters.ClickunderADMFormat})
	}

	// Фильтрация по FraudScores
	if len(filters.FraudScores) > 0 {
		q = q.Where(squirrel.Eq{sspDDL.ColumnFraudScore: filters.FraudScores})
	}

	return q
}
