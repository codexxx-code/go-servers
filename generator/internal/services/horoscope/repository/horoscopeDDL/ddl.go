package horoscopeDDL

import "generator/internal/ddl"

const (
	Table          = ddl.SchemaHoroscope + "." + "horoscopes"
	TableWithAlias = Table + " " + alias
	alias          = "dh"
)

const (
	ColumnID     = "id"
	ColumnDate   = "date"
	ColumnZodiac = "zodiac"
	ColumnText   = "text"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
