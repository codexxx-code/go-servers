package beforeResponseToSSP

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/decimal"
)

type ExchangeService interface {
	CreateAnalyticEventResponseToSSP(ctx context.Context, dto *model.AuctionDTO, exchangeBidID string, sspResponseBid analyticWriterModel.AnalyticResponseBidModel) error
	CreateAnalyticEventResponseToSSPForClickunder(ctx context.Context, dto *model.AuctionDTO, sspResponseBid analyticWriterModel.AnalyticResponseBidModel) error
}

type addPriceForClickunder struct {
	baseLink
	exchangeService ExchangeService
}

func (r *addPriceForClickunder) Apply(dto *beforeResponseToSSP) (err error) {

	// Если цену уже проставили, то ничего не делаем
	if dto.chainSettings.priceAlreadySet {
		return nil
	}

	// Если у нас запрос не на кликандер, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	// Проставляем флаг, что цена уже проставлена
	defer func() {
		dto.chainSettings.priceAlreadySet = true
	}()

	// Собираем сумму цен из всех победивших бидов
	var sumWonBidPricesInSSPCurrency decimal.Decimal

	// Проходимся по каждому победившему биду
	for _, wonBid := range dto.wonBids {

		// Добавляем цену, за которую мы будем биллить DSP (в нее уже вшита маржа на уровне повышения цены бидфлура)
		wonBidPriceInSSPCurrency, err := currencyConverter.Convert(
			wonBid.BillingPriceInDSPCurrency, // Сумма, за которую мы будем биллить DSP
			wonBid.BidResponse.Currency,      // Валюта из победившего бида
			dto.RequestCurrency,              // Валюта запроса
			dto.CurrencyRates,
		)
		if err != nil {
			return err
		}

		// Уменьшаем цену на коэффициент нашей маржи, чтобы SSP забиллила нас по заниженной цене, а мы DSP по ее ставке
		wonBidPriceInSSPCurrency = wonBidPriceInSSPCurrency.Mul(decimal.NewFromInt(1).Sub(dto.Settings.Margin)).Round() // Price = price * (1 - margin)

		sumWonBidPricesInSSPCurrency = sumWonBidPricesInSSPCurrency.Add(wonBidPriceInSSPCurrency)
	}

	// Присваиваем сумму цен из всех ответов DSP в первый бид из ответа SSP
	dto.bidResponse.SeatBids[0].Bids[0].Price = sumWonBidPricesInSSPCurrency

	// Конвертируем цену в дефолтную валюту системы для аналитики
	totalPriceInDefaultCurrency, err := currencyConverter.Convert(
		sumWonBidPricesInSSPCurrency,      // Цена в валюте запроса
		dto.RequestCurrency,               // Валюта запроса
		currencyConverter.DefaultCurrency, // Дефолтная валюта системы
		dto.CurrencyRates,
	)
	if err != nil {
		return err
	}

	// Пишем в аналитику по каждому биду
	if err = r.exchangeService.CreateAnalyticEventResponseToSSPForClickunder(dto.Ctx, dto.AuctionDTO, analyticWriterModel.AnalyticResponseBidModel{
		PriceInDefaultCurrency: totalPriceInDefaultCurrency.String(),
	}); err != nil {
		return err
	}

	return nil
}
