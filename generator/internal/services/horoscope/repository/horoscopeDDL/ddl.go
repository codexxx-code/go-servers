package horoscopeDDL

import "generator/internal/ddl"

const (
	Table          = ddl.SchemaHoroscope + "." + "horoscopes"
	TableWithAlias = Table + " " + alias
	alias          = "dh"
)

const (
	ColumnID              = "id"
	ColumnDateFrom        = "date_from"
	ColumnDateTo          = "date_to"
	ColumnPrimaryZodiac   = "primary_zodiac"
	ColumnSecondaryZodiac = "secondary_zodiac"
	ColumnLanguage        = "language"
	ColumnTimeframe       = "timeframe"
	ColumnHoroscopeType   = "horoscope_type"
	ColumnText            = "text"
)

func WithPrefix(column string) string {
	return alias + "." + column
}
