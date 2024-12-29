package service

import (
	"context"
	"sync/atomic"

	"pkg/openrtb"
)

// SSPBidRequestToExchangeCounter - Переменная, используемая для контроля за количеством сохраняемых запросов в кафку
var SSPBidRequestToExchangeCounter atomic.Int32

// logEachSSPBidRequestToExchange - Переменная, отвечающая за то, каждый какой запрос от SSP сохранять в кафку
const logEachSSPBidRequestToExchange = 500

func (s *EventService) CreateSSPBidRequestToExchangeEvent(ctx context.Context, req openrtb.BidRequest) error {

	// Инкрементируем значение переменной, отвечающей за количество запросов
	SSPBidRequestToExchangeCounter.Add(1)

	// Если это 500 по счету запрос
	if SSPBidRequestToExchangeCounter.Load() == logEachSSPBidRequestToExchange {

		// Обнуляем счетчик
		SSPBidRequestToExchangeCounter.Store(0)

		// Сохраняем запрос в кафку
		return s.eventRepository.CreateSSPBidRequestToExchangeEvent(ctx, req)
	}

	return nil
}
