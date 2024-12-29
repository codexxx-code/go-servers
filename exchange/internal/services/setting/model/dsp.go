package model

import "pkg/decimal"

type Settings struct {
	Margin                     decimal.Decimal `db:"margin"`                         // Маржа системы
	Host                       string          `db:"host"`                           // Хост системы
	DefaultTimeout             int             `db:"default_timeout"`                // Таймаут ответа по умолчанию. В миллисекундах
	EmptySecondPriceReduceCoef decimal.Decimal `db:"empty_second_price_reduce_coef"` // Коэффициент, который мы отнимаем от цены победившей DSP, если она единственная, кто ответил
	ReduceTimeoutCoef          float64         `db:"reduce_timeout_coef"`            // Коэффициент, который мы отнимаем от таймаута SSP для таймаута DSP
	ShowcaseURL                string          `db:"showcase_url"`                   // URL витрины попандер/кликандер трафика
}
