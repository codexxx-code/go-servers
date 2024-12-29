package model

import "pkg/decimal"

// AnalyticResponseBidModel - Основная информация из одного бида ответа (от DSP или от нас)
type AnalyticResponseBidModel struct {
	PriceInDefaultCurrency string `json:"price_in_default_currency" bson:"price_in_default_currency"` // Цена в ответе
}

func GetClearAnalyticResponseBidModel() AnalyticResponseBidModel {
	return AnalyticResponseBidModel{
		PriceInDefaultCurrency: decimal.Zero.String(),
	}
}
