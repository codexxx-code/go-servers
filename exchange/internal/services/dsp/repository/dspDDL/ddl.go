package dspDDL

import "exchange/internal/ddl"

const (
	Table          = ddl.SchemaSSP + "." + "dsps"
	TableWithAlias = Table + " " + alias
	alias          = "d"
)

const (
	ColumnSlug                     = "slug"
	ColumnName                     = "name"
	ColumnIsSupportMultiimpression = "is_support_multiimpression"
	ColumnUrl                      = "url"
	ColumnCurrency                 = "currency"
	ColumnAuctionSecondPrice       = "auction_second_price"
	ColumnBillingUrlType           = "billing_url_type"
	ColumnIsEnable                 = "is_enable"
	ColumnIntegrationType          = "integration_type"
	ColumnSourceTrafficTypes       = "source_traffic_types"
	ColumnFormatTypes              = "format_types"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
