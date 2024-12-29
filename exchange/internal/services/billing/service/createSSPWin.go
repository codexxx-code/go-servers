package service

import (
	"context"
	"time"

	billingRepository "exchange/internal/services/billing/repository"
	exchangeModel "exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/decimal"
)

func (s *BillingService) createSSPWin(
	ctx context.Context,
	priceFromSSP decimal.Decimal,
	dspResponse exchangeModel.DSPResponse,
	bidAnalytic exchangeModel.AnalyticDTO,
) []error {

	var errs []error

	// Получаем фактическую цену за 1 показ от SSP
	priceFor1ShowFromSSPInSSPCurrency := priceFromSSP.Div(cpmFactor)

	// Конвертируем цену, которую мы должны заплатить за 1 показ в базовую валюту
	priceFor1ShowFromSSPInBaseCurrency := currencyConverter.ConvertWithCoefficient(
		priceFor1ShowFromSSPInSSPCurrency,
		dspResponse.CurrencySSPCoefficient,
	)

	// Добавляем ивент победы SSP по старой технике
	if err := s.billingRepository.CreateSSPWinEvent(ctx, billingRepository.CreateSSPEventWinReq{
		RequestID: dspResponse.ExchangeID,
		Price:     priceFor1ShowFromSSPInBaseCurrency.String(),
		SSP:       dspResponse.SlugSSP,
		Timestamp: time.Time{}, // Заполняется в репозитории
	}); err != nil {
		errs = append(errs, err)
	}

	// Добавляем ивент победы SSP по новой технике
	bidAnalytic.FactPriceSSPInDefaultCurrency = priceFor1ShowFromSSPInBaseCurrency.String()
	if err := s.analyticWriterService.CreateSSPBillingEventReq(ctx, bidAnalytic.ConvertToSSPBilling()); err != nil {
		errs = append(errs, err)
	}

	return errs
}
