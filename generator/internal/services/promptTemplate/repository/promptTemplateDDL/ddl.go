package promptTemplateDDL

import "generator/internal/ddl"

const (
	Table          = ddl.SchemaGenerator + "." + "prompt_templates"
	TableWithAlias = Table + " " + alias
	alias          = "p"
)

const (
	ColumnCase = `"case"`
	ColumnText = "template"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
