package service

import (
	"context"

	"pkg/decimal"
)

func (s *CurrencyService) GetRates(ctx context.Context) (map[string]decimal.Decimal, error) {

	// Получаем все валюты
	currencies, err := s.currencyRepository.GetCurrencies(ctx)
	if err != nil {
		return nil, err
	}

	// Формируем мапу название - курс
	currenciesMap := make(map[string]decimal.Decimal, len(currencies))
	for _, currency := range currencies {
		currenciesMap[currency.Slug] = currency.Rate
	}

	return currenciesMap, nil
}
