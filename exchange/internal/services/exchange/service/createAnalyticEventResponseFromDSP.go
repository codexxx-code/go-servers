package service

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"exchange/internal/services/exchange/model"
	"pkg/currencyConverter"
	"pkg/errors"
)

func (s *ExchangeService) createAnalyticEventResponseFromDSP(ctx context.Context, dto *model.AuctionDTO, wonBids []model.BidPointer) error {

	// Проходимся по каждому биду из ответа
	for _, wonBid := range wonBids {

		// Получаем данные из DTO аналитики по импрешенам
		analyticDTOForImpression, ok := dto.AnalyticDTOByDSPRequestByImpression.Get(wonBid.DSPSlug, wonBid.GetBid().ImpID)
		if !ok {
			return errors.InternalServer.New("Не смогли найти аналитику по импрешену")
		}

		// Конвертируем цену в валюту системы для аналитики
		priceInDefaultCurrency, err := currencyConverter.Convert(
			wonBid.GetBid().Price,             // Цена из победившего бида
			wonBid.BidResponse.Currency,       // Валюта из победившего бида
			currencyConverter.DefaultCurrency, // Дефолтная валюта системы
			dto.CurrencyRates,
		)
		if err != nil {
			return err
		}

		analyticDTOForImpression.DSPResponseBid = analyticWriterModel.AnalyticResponseBidModel{
			PriceInDefaultCurrency: priceInDefaultCurrency.String(),
		}
		analyticDTOForImpression.ExchangeBidID = wonBid.ExchangeBidID

		// Добавялем данные в DTO аналитики по бидам
		dto.AnalyticDTOByBid[wonBid.ExchangeBidID] = analyticDTOForImpression

		// Сохраняем информацию о запросе в аналитику
		if err = s.analyticWriterService.CreateDSPToExchangeResponseEvent(ctx, analyticDTOForImpression.ConvertToDSPResponse()); err != nil {
			return err
		}
	}

	return nil
}
