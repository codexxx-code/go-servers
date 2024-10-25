package promptDDL

import "server/internal/ddl"

const (
	Table          = ddl.SchemaZodiac + "." + "prompts"
	TableWithAlias = Table + " " + alias
	alias          = "p"
)

const (
	ColumnID       = "id"
	ColumnCase     = `"case"`
	ColumnLanguage = "language"
	ColumnText     = "text"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
