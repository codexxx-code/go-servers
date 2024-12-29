package beforeResponseToSSP

import (
	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"pkg/currencyConverter"
	"pkg/decimal"
)

type addPrice struct {
	baseLink
	exchangeService ExchangeService
}

func (r *addPrice) Apply(dto *beforeResponseToSSP) (err error) {

	// Если цену уже проставили, то ничего не делаем
	if dto.chainSettings.priceAlreadySet {
		return nil
	}

	// Проставляем флаг, что цена уже проставлена
	defer func() {
		dto.chainSettings.priceAlreadySet = true
	}()

	// Индикатор, какой бид мы сейчас заполняем в ответе
	currentSeatBidIndex := 0

	// Проходимся по каждому победившему биду
	for _, wonBid := range dto.wonBids {

		// Ставим цену, за которую мы будем биллить DSP
		priceInSSPCurrency, err := currencyConverter.Convert(
			wonBid.BillingPriceInDSPCurrency, // Ставка победившего бида
			wonBid.BidResponse.Currency,      // Валюта из победившего бида
			dto.RequestCurrency,              // Валюта запроса
			dto.CurrencyRates,
		)
		if err != nil {
			return err
		}

		// Уменьшаем цену на коэффициент нашей маржи, чтобы SSP забиллила нас по заниженной цене, а мы DSP по ее ставке
		totalPrice := priceInSSPCurrency.Mul(decimal.NewFromInt(1).Sub(dto.Settings.Margin)).Round() // Price = price * (1 - margin)

		dto.bidResponse.SeatBids[currentSeatBidIndex].Bids[0].Price = totalPrice

		// Конвертируем цену в дефолтную валюту системы
		totalPriceInDefaultCurrency, err := currencyConverter.Convert(
			totalPrice,                        // Цена в валюте запроса
			dto.RequestCurrency,               // Валюта запроса
			currencyConverter.DefaultCurrency, // Дефолтная валюта системы
			dto.CurrencyRates,
		)
		if err != nil {
			return err
		}

		// Пишем в аналитику по каждому биду
		if err = r.exchangeService.CreateAnalyticEventResponseToSSP(dto.Ctx, dto.AuctionDTO, wonBid.ExchangeBidID, analyticWriterModel.AnalyticResponseBidModel{
			PriceInDefaultCurrency: totalPriceInDefaultCurrency.String(),
		}); err != nil {
			return err
		}

		currentSeatBidIndex++
	}

	return nil
}
