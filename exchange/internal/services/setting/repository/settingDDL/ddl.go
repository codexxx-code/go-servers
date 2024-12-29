package settingDDL

import "exchange/internal/ddl"

const (
	Table          = ddl.SchemaSSP + "." + "settings"
	TableWithAlias = Table + " " + alias
	alias          = "st"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
