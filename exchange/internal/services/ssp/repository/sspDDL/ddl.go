package sspDDL

import "exchange/internal/ddl"

const (
	Table          = ddl.SchemaSSP + "." + "ssps"
	TableWithAlias = Table + " " + alias
	alias          = "sp"
)

const (
	ColumnSlug                = "slug"
	ColumnName                = "name"
	ColumnTimeout             = "timeout"
	ColumnIsEnable            = "is_enable"
	ColumnIntegrationType     = "integration_type"
	ColumnBillingType         = "billing_type"
	ColumnAuctionSecondPrice  = "auction_second_price"
	ColumnCurrency            = "currency"
	ColumnFraudScore          = "fraud_score"
	ColumnSourceTrafficTypes  = "source_traffic_types"
	ColumnFormatTypes         = "format_types"
	ColumnClickunderADMFormat = "clickunder_adm_format"
	ColumnClickunderDrumSize  = "clickunder_drum_size"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
