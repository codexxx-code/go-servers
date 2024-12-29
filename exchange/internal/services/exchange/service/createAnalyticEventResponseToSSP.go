package service

import (
	"context"

	analyticWriterModel "exchange/internal/services/analyticWriter/model"
	"exchange/internal/services/exchange/model"
	"pkg/errors"
)

func (s *ExchangeService) CreateAnalyticEventResponseToSSP(
	ctx context.Context,
	dto *model.AuctionDTO,
	exchangeBidID string,
	sspResponseBid analyticWriterModel.AnalyticResponseBidModel,
) error {

	// Получаем аналитику по истории этого бида
	analyticDTOForBid, ok := dto.AnalyticDTOByBid[exchangeBidID]
	if !ok {
		return errors.InternalServer.New("Не смогли найти аналитику по биду")
	}

	analyticDTOForBid.SSPResponseBid = sspResponseBid

	dto.AnalyticDTOByBid[exchangeBidID] = analyticDTOForBid

	// Сохраняем информацию о запросе в аналитику
	if err := s.analyticWriterService.CreateExchangeToSSPResponseEvent(ctx, analyticDTOForBid.ConvertToSSPResponse()); err != nil {
		return err
	}

	return nil
}
