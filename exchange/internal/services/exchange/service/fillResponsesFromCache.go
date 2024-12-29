package service

import (
	"context"

	"exchange/internal/services/exchange/model"
	"pkg/errors"
)

func (s *ExchangeService) fillBidsFromCache(dto *model.AuctionDTO, bidsFromDSP []model.BidPointer) (_ []model.BidPointer, err error) {

	requiredBidsCount := len(dto.BidRequest.Impressions)
	haveBidsCount := len(bidsFromDSP)
	needFillBidsCount := requiredBidsCount - haveBidsCount

	// Получаем необходимое количество неиспользуемых ответов
	unusedResponses, err := s.exchangeRepository.PopUnusedBids(dto.Ctx, needFillBidsCount)
	if err != nil {
		return nil, err
	}

	// Если не нашлось нужное количество неиспользованных ответов
	if len(unusedResponses) < needFillBidsCount {

		// Сохраняем неиспользованные ответы этого запроса
		if err = s.exchangeRepository.SaveUnusedBids(context.TODO(), bidsFromDSP); err != nil {
			return nil, err
		}

		return nil, errors.BadRequest.New("Не смогли набрать барабан из ответов от DSP и кэша", errors.ParamsOption(
			"requiredBidsCount", requiredBidsCount,
			"haveBidsCount", haveBidsCount,
			"unusedResponsesCount", len(unusedResponses),
		))
	}

	// Добавляем неиспользованные ответы в список бидов
	bidsFromDSP = append(bidsFromDSP, unusedResponses...)

	return bidsFromDSP, nil
}
