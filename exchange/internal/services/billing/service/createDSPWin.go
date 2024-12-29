package service

import (
	"context"
	"time"

	billingRepository "exchange/internal/services/billing/repository"
	exchangeModel "exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/decimal"
)

func (s *BillingService) createDSPWin(ctx context.Context, dspResponse exchangeModel.DSPResponse, sspFactPrice decimal.Decimal, analyticDTO exchangeModel.AnalyticDTO) []error {

	var errs []error

	// Получаем цену одного показа для DSP в валюте DSP, так как на этапе аукциона мы сохранили цену победы DSP за 1000 показов
	priceFor1ShowForDSPInDSPCurrency := dspResponse.BillingPriceInDSPCurrency.Div(cpmFactor)

	// Конвертируем цену одного показа для DSP в базовую валюту для статистики
	priceFor1ShowForDSPInBaseCurrency := currencyConverter.ConvertWithCoefficient(
		priceFor1ShowForDSPInDSPCurrency,
		dspResponse.CurrencyDSPCoefficient,
	)

	// Сохраняем ивент победы DSP по старой технике
	if err := s.billingRepository.CreateDSPWinEvent(ctx, billingRepository.CreateDSPEventWinReq{
		RequestID: dspResponse.ExchangeBidID, // Тут указываем BidID, потому что для нескольких записей может использоваться один и тот же ExchangeID
		Price:     priceFor1ShowForDSPInBaseCurrency.String(),
		DSP:       dspResponse.SlugDSP,
		Timestamp: time.Time{}, // Заполняется в репозитории
	}); err != nil {
		errs = append(errs, err)
	}

	var drumSize int32 = 1
	if dspResponse.DrumSize != nil {
		drumSize = *dspResponse.DrumSize
	}

	// Сохраняем ивент победы DSP по новой технике

	// Получаем фактическую цену за 1 показ от SSP
	priceFor1ShowFromSSPInSSPCurrency := sspFactPrice.Div(cpmFactor)

	// Конвертируем цену, которую мы должны заплатить за 1 показ в базовую валюту
	priceFor1ShowFromSSPInBaseCurrency := currencyConverter.ConvertWithCoefficient(
		priceFor1ShowFromSSPInSSPCurrency,
		dspResponse.CurrencySSPCoefficient,
	)

	analyticDTO.FactPriceDSPInDefaultCurrency = priceFor1ShowForDSPInBaseCurrency.String()
	analyticDTO.FactPriceSSPInDefaultCurrency = priceFor1ShowFromSSPInBaseCurrency.Div(decimal.NewFromInt(int(drumSize))).String() // Делим цену SSP на размер барабана

	// Получаем цену, которую мы отправили в SSP и делим ее на размер барабана
	sspPrice, err := decimal.NewFromString(analyticDTO.SSPResponseBid.PriceInDefaultCurrency)
	if err != nil {
		errs = append(errs, err)
		return errs
	}
	analyticDTO.SSPResponseBid.PriceInDefaultCurrency = sspPrice.Div(decimal.NewFromInt(int(drumSize))).String() // Делим цену SSP на размер барабана

	if err := s.analyticWriterService.CreateDSPBillingEventReq(ctx, analyticDTO.ConvertToDSPBilling()); err != nil {
		errs = append(errs, err)
	}

	return errs
}
