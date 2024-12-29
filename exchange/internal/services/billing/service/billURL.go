package service

import (
	"context"
	"sync"

	"exchange/internal/metrics"
	"exchange/internal/services/billing/model"
	"exchange/internal/services/billing/service/utils"
	dspModel "exchange/internal/services/dsp/model"
	exchangeModel "exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/model/idType"
	"pkg/decimal"
	"pkg/errors"
	"pkg/log"
	"pkg/slices"
)

// Делим полученную цену на 1000, чтобы получить цену одного показа для статистики, т.к. SSP отправляет стоимость за 1000 показов
var cpmFactor = decimal.NewFromInt(1000) // nolint:gomnd

// BillURL фиксирует нашу победу у SSP и биллит DSP
func (s *BillingService) BillURL(ctx context.Context, req model.BillURLReq) []error {

	// Вызываем биллинг
	errs := s.billURL(ctx, req)

	// Если есть ошибки
	if len(errs) != 0 {

		// Логируем ошибки
		for _, err := range errs {
			log.Error(ctx, err)
		}

		// TODO: Мы тут должны останавливать торги, так как криво пишем аналитику
	}

	return errs
}

func (s *BillingService) billURL(ctx context.Context, req model.BillURLReq) []error {
	var errs []error

	// Получаем цену из входных данных
	priceFromSSP, err := utils.GetPriceFromURL(req)
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	var (
		dspResponses []exchangeModel.DSPResponse
		analyticDTOs []exchangeModel.AnalyticDTO
		isClickunder bool
	)

	switch req.IDType {
	case idType.ExchangeID: // Если передан идентификатор запроса
		isClickunder = true

		// Получаем все связанные биды, которые были собраны в барабан
		dspResponses, err = s.exchangeRepository.GetDSPResponses(ctx, exchangeModel.GetDSPResponsesReq{ //nolint:exhaustruct
			RequestIDs: []string{req.ID},
		})

		// Получаем весь путь запроса для этого биллинга
		analyticDTOs, err = s.exchangeRepository.GetAnalyticDTOs(ctx, exchangeModel.GetAnalyticDTOsReq{ //nolint:exhaustruct
			RequestIDs: []string{req.ID},
		})

	case idType.ExchangeBidID: // Если передан идентификатор бида
		isClickunder = false

		// Получаем победивший бид
		dspResponses, err = s.exchangeRepository.GetDSPResponses(ctx, exchangeModel.GetDSPResponsesReq{ //nolint:exhaustruct
			BidIDs: []string{req.ID},
		})

		// Получаем весь путь запроса для этого биллинга
		analyticDTOs, err = s.exchangeRepository.GetAnalyticDTOs(ctx, exchangeModel.GetAnalyticDTOsReq{ //nolint:exhaustruct
			BidIDs: []string{req.ID},
		})
	}
	if err != nil {
		errs = append(errs, err)
		return errs
	}

	if len(dspResponses) == 0 {
		errs = append(errs, errors.InternalServer.New("Не найдены ответы от DSP для биллинга"))
		return errs
	}
	metrics.IncEventCall("win", dspResponses[0].SlugSSP, dspResponses[0].RequestPublisherID)

	if len(analyticDTOs) == 0 {
		errs = append(errs, errors.InternalServer.New("Не найдена аналитика для биллинга"))
		return errs
	}

	analyticDTOsByExchangeBidIDMap := slices.ToMap(analyticDTOs, func(analyticDTO exchangeModel.AnalyticDTO) string {
		return analyticDTO.ExchangeBidID
	})

	// Создаем ивент победы SSP
	if isClickunder {

		// Берем рандомную аналитику, потому что в каждой из них есть нужная информация о SSP, вся информация о DSP при аггрегации теряется
		if _errs := s.createSSPWinForClickunder(ctx, priceFromSSP, dspResponses, analyticDTOs[0]); _errs != nil {
			errs = append(errs, _errs...)
		}
	} else {
		if _errs := s.createSSPWin(ctx, priceFromSSP, dspResponses[0], analyticDTOs[0]); len(_errs) != 0 {
			errs = append(errs, _errs...)
		}
	}

	// Собираем слаги DSP
	dspSlugs := slices.GetUniqueFields(dspResponses, func(dspResponse exchangeModel.DSPResponse) string {
		return dspResponse.SlugDSP
	})

	// Получаем все необходимые DSP по слагам
	dsps, err := s.dspService.GetDSPs(ctx, dspModel.GetDSPsReq{ //nolint:exhaustruct
		Slugs: dspSlugs,
	})
	if err != nil {
		errs = append(errs, err)
	}

	// Делаем мапу для быстрого доступа к DSP по слагу
	dspsMap := slices.ToMap(dsps, func(dsp dspModel.DSP) string {
		return dsp.Slug
	})

	wg := &sync.WaitGroup{}
	wg.Add(len(dspResponses))

	// Проходимся по каждому биду
	for _, dspResponse := range dspResponses {

		dsp, ok := dspsMap[dspResponse.SlugDSP]
		if !ok {
			errs = append(errs, errors.InternalServer.New("DSP не найдена в системе", errors.ParamsOption(
				"slug", dsp.Slug,
			)))
			continue
		}

		// Биллим DSP
		go func() {
			if err = s.billDSP(ctx, dsp, dspResponse); err != nil {
				errs = append(errs, err)
			}
			wg.Done()
		}()

		// Получаем аналитику для этого бида
		analyticDTO, ok := analyticDTOsByExchangeBidIDMap[dspResponse.ExchangeBidID]
		if !ok {
			errs = append(errs, errors.InternalServer.New("Аналитика для бида не найдена", errors.ParamsOption(
				"exchangeBidID", dspResponse.ExchangeBidID,
			)))
			continue
		}

		// Создаем ивент победы DSP
		if _errs := s.createDSPWin(ctx, dspResponse, priceFromSSP, analyticDTO); len(_errs) != 0 {
			errs = append(errs, _errs...)
		}
	}

	wg.Wait()

	return errs
}
