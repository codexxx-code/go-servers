package service

import (
	"sync"

	"exchange/internal/metrics"
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/service/chain/beforeRequestToDSP"
	"exchange/internal/services/exchange/service/utils"
	"pkg/log"
	"pkg/openrtb"
)

func (s *ExchangeService) sendRequestToDSPAsync(
	wg *sync.WaitGroup,
	chRes chan SendRequestToDSPDTO,
	bidRequest openrtb.BidRequest,
	dto *model.AuctionDTO,
	dsp dspModel.DSP,
	requestNumber int,
) {
	defer wg.Done()

	// Модернизируем запрос для конкретной DSP
	var err error
	if bidRequest, err = beforeRequestToDSP.RunChain(dto, bidRequest, dsp); err != nil {
		chRes <- SendRequestToDSPDTO{
			BidResponse: nil,
			DSP:         dsp,
			Err:         err,
		}
		return
	}

	// Логгируем запрос в DSP
	go func() {

		// Добавляем к запросу метаданные об SSP и DSP
		rtbRequestWithChangedExt, err := utils.FillExtField(bidRequest, map[string]any{
			"ssp-slug": dto.SSP.Slug,
			"dsp-slug": dsp.Slug,
		})
		if err != nil {
			log.Error(dto.Ctx, err)
		}

		// Сохраняем запрос в кафку
		if err := s.eventService.CreateExchangeBidRequestToDSPEvent(dto.Ctx, requestNumber, rtbRequestWithChangedExt); err != nil {
			log.Error(dto.Ctx, err)
		}
	}()

	if err := s.createAnalyticEventRequestToDSP(dto.Ctx, dto, bidRequest, dsp); err != nil {
		chRes <- SendRequestToDSPDTO{
			BidResponse: nil,
			DSP:         dsp,
			Err:         err,
		}
	}

	// Отправляем запрос в DSP
	bidResponseFromDSP, statusCode, err := s.exchangeNetwork.SendBidRequestToDSP(dto.Ctx, requestNumber, dsp.URL, bidRequest) // Используем везде только копию, это важно!

	// Фиксируем код ответа в метриках
	metrics.IncStatusCodeFromDSPRTB(statusCode, dsp.Slug, dto.SSP.Slug)

	if err != nil {
		chRes <- SendRequestToDSPDTO{
			BidResponse: nil,
			DSP:         dsp,
			Err:         err,
		}
		return
	}

	// Отправляем ответ в канал
	chRes <- SendRequestToDSPDTO{
		BidResponse: &bidResponseFromDSP,
		DSP:         dsp,
		Err:         nil,
	}
}
