package service

import (
	"context"
	"time"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	billingRepository "exchange/internal/services/billing/repository"
	exchangeModel "exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/decimal"
	"pkg/errors"
)

func (s *BillingService) createSSPWinForClickunder(
	ctx context.Context,
	priceFromSSP decimal.Decimal,
	dspResponses []exchangeModel.DSPResponse,
	analyticDTO exchangeModel.AnalyticDTO,
) []error {

	var errs []error

	var (
		currencySSPCoefficient decimal.Decimal
		slugSSP                string
		requestID              string
	)

	// Проверяем, что все ответы от DSP содержат одинаковые коэффициенты валюты для SSP и слаг SSP, которая послала запрос, а также одинаковый requestID
	for _, dspResponse := range dspResponses {
		switch {
		case currencySSPCoefficient.IsZero():
			currencySSPCoefficient = dspResponse.CurrencySSPCoefficient
		case !currencySSPCoefficient.Equal(dspResponse.CurrencySSPCoefficient):
			errs = append(errs, errors.InternalServer.New("different currency coefficients in dsp responses",
				errors.ParamsOption(
					"currencySSPCoefficient", currencySSPCoefficient.String(),
					"anotherCurrencySSPCoefficient", dspResponse.CurrencySSPCoefficient.String(),
					"requestID", dspResponse.ExchangeID,
				)),
			)
		}
		switch {
		case slugSSP == "":
			slugSSP = dspResponse.SlugSSP
		case slugSSP != dspResponse.SlugSSP:
			errs = append(errs, errors.InternalServer.New("different slugs in dsp responses",
				errors.ParamsOption(
					"slugSSP", slugSSP,
					"anotherSlugSSP", dspResponse.SlugSSP,
					"requestID", dspResponse.ExchangeID,
				)),
			)
		}
		switch {
		case requestID == "":
			requestID = dspResponse.ExchangeID
		case requestID != dspResponse.ExchangeID:
			errs = append(errs, errors.InternalServer.New("different requestIDs in dsp responses",
				errors.ParamsOption(
					"requestID", requestID,
					"anotherRequestID", dspResponse.ExchangeID,
					"requestID", dspResponse.ExchangeID,
				)),
			)
		}
	}

	// Получаем фактическую цену за 1 показ от SSP
	priceFor1ShowFromSSPInSSPCurrency := priceFromSSP.Div(cpmFactor)

	// Конвертируем цену, которую мы должны заплатить за 1 показ в базовую валюту
	priceFor1ShowFromSSPInBaseCurrency := currencyConverter.ConvertWithCoefficient(
		priceFor1ShowFromSSPInSSPCurrency,
		currencySSPCoefficient,
	)

	// Добавляем ивент победы SSP по старой технике
	if err := s.billingRepository.CreateSSPWinEvent(ctx, billingRepository.CreateSSPEventWinReq{
		RequestID: requestID,
		Price:     priceFor1ShowFromSSPInBaseCurrency.String(),
		SSP:       slugSSP,
		Timestamp: time.Time{}, // Заполняется в репозитории
	}); err != nil {
		errs = append(errs, err)
	}

	// Так как кликандер ответ по сути состоит из нескольких ответов от DSP, а нам надо создать один ивент победы SSP,
	// То данные по запросу и ответу DSP мы при аггрегации теряем

	// Добавляем ивент победы SSP по новой технике
	analyticDTO.FactPriceSSPInDefaultCurrency = priceFor1ShowFromSSPInBaseCurrency.String()
	analyticDTO.DSPRequestImpression = analyticWriterModel.GetClearAnalyticRequestImpressionModel()
	analyticDTO.DSPResponseBid = analyticWriterModel.GetClearAnalyticResponseBidModel()
	analyticDTO.FactPriceDSPInDefaultCurrency = decimal.Zero.String()
	if err := s.analyticWriterService.CreateSSPBillingEventReq(ctx, analyticDTO.ConvertToSSPBilling()); err != nil {
		errs = append(errs, err)
	}

	return errs
}
