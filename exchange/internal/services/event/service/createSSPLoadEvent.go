package service

import (
	"context"

	"exchange/internal/metrics"
	"exchange/internal/services/event/model"
	exchangeModel "exchange/internal/services/exchange/model"
	"pkg/errors"
	"pkg/pointer"
	"pkg/slices"
)

func (s *EventService) CreateSSPLoadEvent(ctx context.Context, req model.CreateSSPLoadEventReq) (err error) {

	// Получаем любой бид, который был собран в барабан
	bidResponse, err := slices.FirstWithError(
		s.exchangeRepository.GetDSPResponses(ctx, exchangeModel.GetDSPResponsesReq{ //nolint:exhaustruct
			RequestIDs: []string{req.ID},
			Limit:      pointer.Pointer(int64(1)),
		}),
	)
	if err != nil {
		return err
	}

	metrics.IncEventCall("load", bidResponse.SlugDSP, bidResponse.RequestPublisherID)

	// Проверяем, есть ли идентификатор паблишера
	if bidResponse.RequestPublisherID == "" {
		return errors.InternalServer.New("Не нашли идентификатор паблишера в BidResponses", errors.ParamsOption(
			"id", req.ID,
		))
	}

	// Инкрементируем значение загрузок на паблишера
	return s.eventRepository.IncrementLoadsForPublisher(ctx, bidResponse.RequestPublisherID)
}
