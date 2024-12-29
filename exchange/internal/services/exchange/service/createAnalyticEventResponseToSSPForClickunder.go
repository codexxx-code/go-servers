package service

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"exchange/internal/services/exchange/model"
)

func (s *ExchangeService) CreateAnalyticEventResponseToSSPForClickunder(
	ctx context.Context,
	dto *model.AuctionDTO,
	sspResponseBid analyticWriterModel.AnalyticResponseBidModel,
) error {
	// Так как кликандер ответ состоит из нескольких ответов из DSP, а нам в аналитике надо показать один ответ в SSP (фактически ведь произойдет один ответ),
	// То мы просто берем бид первого запроса из SSP и сохраняем только его, так как другая информация является аггрегированной

	var randomAnalyticDTOForBid model.AnalyticDTO
	isFirst := true

	// Сохраняем информацию об ответе в модель аналитики
	for exchangeBidID, analyticDTO := range dto.AnalyticDTOByBid {
		analyticDTO.SSPResponseBid = sspResponseBid
		dto.AnalyticDTOByBid[exchangeBidID] = analyticDTO
		if isFirst {
			randomAnalyticDTOForBid = dto.AnalyticDTOByBid[exchangeBidID]
			isFirst = false
		}
	}

	// Сбрасываем информацию о запросе в DSP и ответе от DSP для записи в аналитику
	randomAnalyticDTOForBid.DSPRequestImpression = analyticWriterModel.GetClearAnalyticRequestImpressionModel()
	randomAnalyticDTOForBid.DSPResponseBid = analyticWriterModel.GetClearAnalyticResponseBidModel()

	// Сохраняем информацию о запросе в аналитику
	if err := s.analyticWriterService.CreateExchangeToSSPResponseEvent(ctx, randomAnalyticDTOForBid.ConvertToSSPResponse()); err != nil {
		return err
	}

	return nil
}
