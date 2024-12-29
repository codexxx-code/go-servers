package model

import "pkg/decimal"

type Currency struct {
	Slug   string          `db:"slug"`   // Строковый идентификатор валюты в формате ISO-4217
	Name   string          `db:"name"`   // Человекочитаемое название валюты
	Rate   decimal.Decimal `db:"rate"`   // Курс валюты относительно доллара
	Symbol string          `db:"symbol"` // Символ валюты
}
