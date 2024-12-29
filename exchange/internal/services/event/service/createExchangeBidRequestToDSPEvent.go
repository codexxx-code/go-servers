package service

import (
	"context"
	"sync/atomic"

	"pkg/openrtb"
)

// ExchangeBidRequestToDSPCounter - Переменная, используемая для контроля за количеством сохраняемых запросов в кафку
var ExchangeBidRequestToDSPCounter atomic.Int32

// logEachExchangeBidRequestToDSP - Переменная, отвечающая за то, каждый какой запрос в DSP сохранять в кафку
const logEachExchangeBidRequestToDSP = 500

func (s *EventService) CreateExchangeBidRequestToDSPEvent(ctx context.Context, _ int, req openrtb.BidRequest) error {

	// Инкрементируем значение переменной, отвечающей за количество запросов
	ExchangeBidRequestToDSPCounter.Add(1)

	// Если это 500 по счету запрос
	if ExchangeBidRequestToDSPCounter.Load() == logEachExchangeBidRequestToDSP {

		// Обнуляем счетчик
		ExchangeBidRequestToDSPCounter.Store(0)

		// Сохраняем запрос в кафку
		return s.eventRepository.CreateExchangeBidRequestToDSPEvent(ctx, req)
	}

	return nil

}
