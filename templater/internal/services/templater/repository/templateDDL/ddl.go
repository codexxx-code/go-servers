package templateDDL

import "templater/internal/ddl"

const (
	Table          = ddl.SchemaTemplater + "." + "templates"
	TableWithAlias = Table + " " + alias
	alias          = "t"
)

const (
	ColumnID       = "id"
	ColumnSSPSlug  = "ssp_slug"
	ColumnTemplate = "template"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
