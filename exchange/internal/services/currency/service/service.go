package service

import (
	"context"

	"exchange/internal/services/currency/model"
	currencyRepository "exchange/internal/services/currency/repository"
)

type CurrencyService struct {
	currencyRepository CurrencyRepository
}

func NewCurrencyService(
	currencyRepository CurrencyRepository,
) *CurrencyService {
	return &CurrencyService{
		currencyRepository: currencyRepository,
	}
}

var _ CurrencyRepository = &currencyRepository.CurrencyRepository{}

type CurrencyRepository interface {
	GetCurrencies(context.Context) ([]model.Currency, error)
}
