package service

import (
	"sync"

	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
)

type SendRequestToDSPDTO struct {
	BidResponse *openrtb.BidResponse
	DSP         dspModel.DSP
	Err         error
}

func (s *ExchangeService) sendRequestsToAllDSPs(dto *model.AuctionDTO) (results []SendRequestToDSPDTO) {

	wg := &sync.WaitGroup{}

	// Количество запросов, которые мы сделаем
	var countRequests int

	// Проходимся по каждой DSP
	for _, dsp := range dto.DSPs {

		if dsp.IsSupportMultiimpression { // Если DSP поддерживает мультиимпрешены

			// Делаем один запрос
			countRequests++

		} else { // Если DSP не поддерживает мультиимпрешены

			// Делаем свой запрос для каждого импрешена
			countRequests += len(dto.BidRequest.Impressions)
		}
	}

	chRes := make(chan SendRequestToDSPDTO, countRequests)
	wg.Add(countRequests)

	requestNumber := 0

	// Проходимся по каждой DSP
	for _, dsp := range dto.DSPs {

		// Если DSP поддерживает мультиимпрешены
		if dsp.IsSupportMultiimpression {

			// Отправляем мультиимпрешен запрос в горутине
			requestNumber++
			go s.sendRequestToDSPAsync(wg, chRes, dto.BidRequest, dto, dsp, requestNumber)

		} else { // Если DSP не поддерживает мультиимпрешены

			// Проходимся по каждому импрешену
			for _, impression := range dto.BidRequest.Impressions {

				// Копируем весь запрос, чтобы не испортить его
				bidRequestCopy := dto.BidRequest.Copy()

				// Делаем запрос с одним импрешеном
				bidRequestCopy.Impressions = []openrtb.Impression{impression}

				// Отправляем запрос с одним импрешеном в горутине
				requestNumber++
				go s.sendRequestToDSPAsync(wg, chRes, bidRequestCopy, dto, dsp, requestNumber)
			}
		}
	}

	// Ждем ответа от всех DSP - кто не успел за таймаут, тот отвалится с ошибкой
	wg.Wait()
	close(chRes)

	// Собираем результаты
	results = make([]SendRequestToDSPDTO, 0, countRequests)
	for res := range chRes {
		results = append(results, res)
	}

	return results
}
