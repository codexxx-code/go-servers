package sspToDSPDDL

import "exchange/internal/ddl"

const (
	Table          = ddl.SchemaSSP + "." + "ssps_to_dsps"
	TableWithAlias = Table + " " + alias
	alias          = "std"
)

const (
	ColumnSSPSlug   = "ssp_slug"
	ColumnDSPSlug   = "dsp_slug"
	ColumnIsDeleted = "is_deleted"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
