package promptDDL

import "generator/internal/ddl"

const (
	Table          = ddl.SchemaGenerator + "." + "prompts"
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
