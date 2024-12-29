package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"

	"exchange/internal/services/dsp/model/dspFilters"
	"exchange/internal/services/dsp/repository/dspDDL"
	"pkg/ddlHelper"
)

func GetDSPsFilters(filters dspFilters.DSPFilters, q squirrel.SelectBuilder) squirrel.SelectBuilder {
	// Фильтрация по Slugs
	if len(filters.Slugs) != 0 {
		q = q.Where(squirrel.Eq{dspDDL.ColumnSlug: filters.Slugs})
	}

	// Фильтрация по Names
	if len(filters.Names) != 0 {
		q = q.Where(squirrel.Eq{dspDDL.ColumnName: filters.Names})
	}

	// Фильтрация по IsEnable
	if filters.IsEnable != nil {
		q = q.Where(squirrel.Eq{dspDDL.ColumnIsEnable: *filters.IsEnable})
	}

	// Фильтрация по IntegrationTypes
	if len(filters.IntegrationTypes) != 0 {
		q = q.Where(squirrel.Eq{dspDDL.ColumnIntegrationType: filters.IntegrationTypes})
	}

	// Фильтрация по SourceTrafficTypes
	if len(filters.SourceTrafficTypes) != 0 {
		q = q.Where(ddlHelper.PartContains(dspDDL.ColumnSourceTrafficTypes, pq.Array(filters.SourceTrafficTypes)))
	}

	// Фильтрация по BillingTypes
	if len(filters.BillingTypes) != 0 {
		q = q.Where(squirrel.Eq{dspDDL.ColumnBillingUrlType: filters.BillingTypes})
	}

	// Фильтрация по AuctionSecondPrice
	if filters.AuctionSecondPrice != nil {
		q = q.Where(squirrel.Eq{dspDDL.ColumnAuctionSecondPrice: *filters.AuctionSecondPrice})
	}

	// Фильтрация по Currencies
	if len(filters.Currencies) != 0 {
		q = q.Where(squirrel.Eq{dspDDL.ColumnCurrency: filters.Currencies})
	}

	// Фильтрация по FormatTypes
	if len(filters.FormatTypes) != 0 {
		q = q.Where(ddlHelper.PartContains(dspDDL.ColumnFormatTypes, pq.Array(filters.FormatTypes)))
	}

	// Фильтрация по IsSupportMultiimpression
	if filters.IsSupportMultiimpression != nil {
		q = q.Where(squirrel.Eq{dspDDL.ColumnIsSupportMultiimpression: *filters.IsSupportMultiimpression})
	}
	return q
}
