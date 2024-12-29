package currencyDDL

import "exchange/internal/ddl"

const (
	Table          = ddl.SchemaSSP + "." + "currencies"
	TableWithAlias = Table + " " + alias
	alias          = "c"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
