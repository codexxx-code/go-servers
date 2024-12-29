package service

import (
	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/service/utils"
	"pkg/errors"
	"pkg/log"
)

// auction проводит аукцион среди выбранных DSP и выбирает победителя
func (s *ExchangeService) auction(dto *model.AuctionDTO) (wonBids []model.WonBid, err error) {

	// Отравляем запрос в каждую DSP и получаем ответы
	responsesFromDSPs := s.sendRequestsToAllDSPs(dto)

	// Формируем dto для функции выбора победивших ставок
	bidResponsesFromDSPs := make([]model.BidResponseFromDSP, 0, len(responsesFromDSPs))

	var errs []error

	// Проходимся по каждому ответу от DSP
	for _, res := range responsesFromDSPs {

		if res.Err != nil {
			errs = append(errs, res.Err)
			continue
		}

		// Валидируем ответ от DSP
		if err = utils.ValidateResponseFromDSP(res.BidResponse, res.DSP); err != nil {
			log.LogError(dto.Ctx, err)
			continue
		}

		// Добавляем ответ в список для выбора победивших ставок
		bidResponsesFromDSPs = append(bidResponsesFromDSPs, model.BidResponseFromDSP{
			DSP:         res.DSP,
			BidResponse: *res.BidResponse,
		})
	}

	// Логгируем ошибки, полученные при запросах к DSP
	if len(errs) > 0 {
		utils.LogNetworkErrors(errs)
	}

	// Получаем из всех ответов от DSP все биды
	bidsFromDSPs := make([]model.BidPointer, 0, len(bidResponsesFromDSPs))
	for _, bidResponseFromDSP := range bidResponsesFromDSPs {
		bidsFromDSPs = append(bidsFromDSPs, bidResponseFromDSP.ExtractBids(dto.MappingImpressionsByDSPs)...)
	}

	// Пишем аналитику по всем ответам от DSP
	if err = s.createAnalyticEventResponseFromDSP(dto.Ctx, dto, bidsFromDSPs); err != nil {
		return nil, err
	}

	// Проверяем, есть ли хоть один бид от DSP
	if len(bidsFromDSPs) == 0 {
		return nil, errors.NotFound.New("Нет ни одного бида от DSPs")
	}

	// Если тип запроса кликандер
	if dto.IsClickunder {

		// Если количество бидов от всех DSP меньше нужного количества импрешенов
		if len(bidsFromDSPs) < len(dto.BidRequest.Impressions) {

			// Дополняем ответы от DSP другими неиспользуемыми бидами
			if bidsFromDSPs, err = s.fillBidsFromCache(dto, bidsFromDSPs); err != nil {
				return nil, err
			}
		}
	}

	// Проводим аукцион и получаем победившие биды с ценами, по которым будем биллить DSP
	wonBids, err = utils.GetWinBids(
		bidsFromDSPs,
		dto.Settings,
		dto.CurrencyRates,
		dto.BidRequest,
	)
	if err != nil {
		return nil, err
	}

	// Проверяем, что набралось такое количество победивших бидов, сколько у нас импрешенов в запросе
	if len(wonBids) != len(dto.BidRequest.Impressions) {
		return nil, errors.BadRequest.New("Недостаточно бидов от DSP", errors.ParamsOption(
			"impressions", len(dto.BidRequest.Impressions),
			"bids", len(wonBids),
		))
	}

	return wonBids, nil
}
