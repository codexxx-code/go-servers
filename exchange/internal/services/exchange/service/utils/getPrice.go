package utils

import (
	"sort"

	"exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/decimal"
	"pkg/errors"
	"pkg/openrtb"

	settingsModel "exchange/internal/services/setting/model"
)

type priceMatrixObject struct {
	model.BidPointer

	PriceInDefaultCurrency decimal.Decimal
}

// GetWinBids возвращает биды и цены, по которым мы будем биллить DSP за эти биды
func GetWinBids(
	bidPointers []model.BidPointer,
	settings settingsModel.Settings,
	currencyRates map[string]decimal.Decimal,
	bidRequest openrtb.BidRequest,
) (res []model.WonBid, err error) {

	// Если нет бидов от DSP, то и победителей нет
	if len(bidPointers) == 0 {
		return nil, errors.BadRequest.New("Нет бидов от DSP")
	}

	// Формируем ценовую матрицу для каждого импрешена impressionID - ставки
	priceMatrix := make(map[string][]priceMatrixObject)

	// Проходимся по каждому биду от DSP
	for _, bidPointer := range bidPointers {

		// Получаем бид из указателя
		bid := bidPointer.GetBid()

		// Получаем цену из ответа DSP в валюте DSP
		priceInDSPCurrency := bid.Price

		// Переводим цену из валюты DSP в базовую валюту (RUB)
		priceInBaseCurrency, err := currencyConverter.Convert(
			priceInDSPCurrency,
			bidPointer.BidResponse.Currency,
			currencyConverter.DefaultCurrency,
			currencyRates,
		)
		if err != nil {
			return nil, err
		}

		// Добавляем эту ставку в массив ставок для этого impressionID
		priceMatrix[bidPointer.ExchangeImpressionID] = append(priceMatrix[bidPointer.ExchangeImpressionID], priceMatrixObject{
			BidPointer:             bidPointer,
			PriceInDefaultCurrency: priceInBaseCurrency,
		})
	}

	// Сортируем ценовую матрицу по убыванию цены для каждого импрешена
	for _, bids := range priceMatrix {
		sort.Slice(bids, func(i, j int) bool {
			return bids[i].PriceInDefaultCurrency.GreaterThanOrEqual(bids[j].PriceInDefaultCurrency)
		})
	}

	// Составляем список победивших ставок с сортировкой по положению импрешена в запросе
	wonBids := make([]model.WonBid, 0, len(bidRequest.Impressions))

	// Проходимся по каждому импрешену
	for _, impression := range bidRequest.Impressions {

		// Получаем все ставки для этого импрешена
		bids := priceMatrix[impression.ID]

		// Проверяем, что для импрешена есть хоть одна ставка
		if len(bids) == 0 {
			return nil, errors.BadRequest.New("Нет ставок для импрешена", errors.ParamsOption(
				"requestID", bidRequest.ID,
				"impressionID", impression.ID,
			))
		}

		// Получаем лучшую ставку
		bestBid := bids[0]

		// Предзаполняем структуру победившей ставки
		wonBid := model.WonBid{
			BidPointer:                bestBid.BidPointer,
			BillingPriceInDSPCurrency: decimal.Zero,
		}

		// TODO: Сделать проверку, что полученная цена не меньше нашего бидфлура

		// Выбираем цену, на которую будем биллить победившую DSP в случае нашей победы
		if !bestBid.IsAuctionSecondPrice { // Если аукцион первой цены

			// Просто проставляем цену победившей DSP
			wonBid.BillingPriceInDSPCurrency = bestBid.PriceInDefaultCurrency

		} else { // Если аукцион второй цены

			// Если есть еще ставки на этот импрешен
			if len(bids) > 1 {

				// Проставляем цену следующего ответа
				wonBid.BillingPriceInDSPCurrency = bids[1].PriceInDefaultCurrency

			} else { // Если больше ответов нет

				// Отнимаем от лучшей ставки коэффициент из настроек // result BillingPriceInDSPCurrency = BillingPriceInDSPCurrency * (1 - reduce coef)
				wonBid.BillingPriceInDSPCurrency = bestBid.PriceInDefaultCurrency.Mul(decimal.NewFromInt(1).Sub(settings.EmptySecondPriceReduceCoef))
			}
		}

		// Конвертируем ставку бида обратно в валюту DSP
		wonBid.BillingPriceInDSPCurrency, err = currencyConverter.Convert(
			wonBid.BillingPriceInDSPCurrency,
			currencyConverter.DefaultCurrency,
			bestBid.BidResponse.Currency,
			currencyRates,
		)
		if err != nil {
			return nil, err
		}

		// Добавляем победившую ставку в список победивших ставок
		wonBids = append(wonBids, wonBid)

		// Удаляем из ценовой матрицы ставку для этого импрешена
		priceMatrix[impression.ID] = bids[1:]
	}

	return wonBids, nil
}
