package service

import (
	"context"
	"fmt"
	"time"

	"exchange/internal/services/exchange/model"
	"exchange/internal/services/exchange/service/chain/beforeAnalytic"
	"exchange/internal/services/exchange/service/chain/beforeAuction"
	"exchange/internal/services/exchange/service/chain/beforeResponseToSSP"
	"exchange/internal/services/exchange/service/utils"
	settingsModel "exchange/internal/services/setting/model"
	"pkg/errors"
	"pkg/log"
	"pkg/maps"
	"pkg/openrtb"
	"pkg/uuid"
)

func (s *ExchangeService) BidSSP(ctx context.Context, req model.SSPBidReq) (res openrtb.BidResponse, err error) {

	// Запоминаем время начала входа в функцию
	timeStart := time.Now()

	// Делаем глубокое копирование исходного запроса, эту копию будем использовать ДЛЯ ВСЕХ ТЕХ ДЕЙСТВИЙ, которые как-либо
	// изменяют запрос (даже казалось бы копированием) потому что иначе мы меняем исходный запрос (даже если передаем копию,
	// так как слайсы в этом объекте все равно остаются ссылками), что недопустимо, так как дальше мы ссылаемся на данные
	// в исходном запросе, они на протяжении всей этой функции должны оставаться неизменными.
	// Соответственно req.BidRequest МЫ НЕ ДОЛЖНЫ НИКАК ИЗМЕНЯТЬ
	bidRequestCopy := req.BidRequest.Copy()

	// Сохраняем запрос в кафку, чтобы потом проанализировать

	// Добавляем к запросу метаданные об SSP
	rtbRequestWithSSPExt, err := utils.FillExtField(req.BidRequest, map[string]any{"ssp-slug": req.SSPSlug})
	if err != nil {
		log.Error(ctx, err)
	}

	// Сохраняем запрос в кафку
	if err := s.eventService.CreateSSPBidRequestToExchangeEvent(ctx, rtbRequestWithSSPExt); err != nil {
		log.Error(ctx, err)
	}

	// // Подготовка входных данных для аукциона

	// Получаем SSP
	ssp, err := s.getSSP(ctx, req.SSPSlug, req.BidRequest)
	if err != nil {
		return res, err
	}

	// Определяем DTO, который будет нас сопровождать на протяжении всего оставшегося пути
	dto := &model.AuctionDTO{
		Ctx:                                 ctx,
		OriginalBidRequest:                  req.BidRequest,
		BidRequest:                          bidRequestCopy,
		SSP:                                 ssp,
		Settings:                            settingsModel.Settings{},                                            // Заполняется позже
		DSPs:                                nil,                                                                 // Заполнится после чейна
		CurrencyRates:                       nil,                                                                 // Заполняется позже
		IsClickunder:                        false,                                                               // Конфигурируется в чейне setIsClickunder.go
		DrumSize:                            nil,                                                                 // Проставляется в чейне setIsClickunder.go
		RequestCurrency:                     "",                                                                  // Проставляется в чейне setRequestCurrency.go
		SourceTrafficType:                   "",                                                                  // Проставляется в чейне setSourceTrafficType.go
		RequestTimeout:                      0,                                                                   // Заполняется позже
		ExchangeID:                          uuid.New(),                                                          // Генерируем новый уникальный идентификатор для всего запроса
		MappingExchangeImpressionIDs:        make(map[string]string, len(bidRequestCopy.Impressions)),            // Заполняется в чейне makeMappingImpressionIDs.go
		MappingImpressionsByDSPs:            model.NewMappingImpressionsByDSPsMu(),                               // Заполняется в чейне changeImpressionIDs.go перед каждым запросом в DSP
		PublisherID:                         "",                                                                  // Заполняется в чейне setPublisherID.go
		GeoCountry:                          "",                                                                  // Заполняется в чейне logGetByIP.go
		AnalyticDTOByImpression:             make(map[string]model.AnalyticDTO, len(bidRequestCopy.Impressions)), // Заполняем в функции createAnalyticEventRequestFromSSP
		AnalyticDTOByDSPRequestByImpression: model.NewAnalyticDTOByDSPRequestByImpressionMu(),                    // Заполняем в функции sendRequestToDSPAsync перед каждым запросом в DSP
		AnalyticDTOByBid:                    make(map[string]model.AnalyticDTO),                                  // Заполняем в функции createAnalyticEventResponseFromDSP
	}

	// Подготавливаем запрос перед отправкой в аналитику
	if err = beforeAnalytic.RunChain(s.exchangeRepository, dto); err != nil {
		return res, err
	}

	// Получаем актуальные курсы валют, которые будут потом сопровождать нас по всему циклу жизни запроса, в т.ч. биллинге
	dto.CurrencyRates, err = s.currencyService.GetRates(ctx)
	if err != nil {
		return res, err
	}

	// Получаем сжатую информацию по запросу в разрезе импррешенов
	if err = s.createAnalyticEventRequestFromSSP(ctx, dto); err != nil {
		return res, err
	}

	// Проверяем, включена ли SSP
	if !ssp.IsEnable {
		return res, errors.BadRequest.New(fmt.Sprintf("SSPRequestImpression \"%s\" is not enabled", ssp.Slug),
			errors.LogAsOption(errors.LogAsDebug),
		)
	}

	// Получаем настройки системы
	dto.Settings, err = s.settingsService.GetSettings(ctx)
	if err != nil {
		return res, err
	}

	// Определяем таймаут запроса и ставим его
	timeout := time.Duration(dto.Settings.DefaultTimeout) * time.Millisecond
	switch {
	case bidRequestCopy.TimeMax != 0:
		timeout = time.Duration(bidRequestCopy.TimeMax) * time.Millisecond
	case ssp.Timeout != nil:
		timeout = time.Duration(*ssp.Timeout) * time.Millisecond
	}
	dto.RequestTimeout = timeout

	// Добавляем ко времени старта время таймаута и обновляем контекст с дедлайном
	ctx, cancel := context.WithDeadline(ctx, timeStart.Add(timeout))
	defer cancel()

	// Меняем запрос перед аукционом (модернизация запроса, касающаяся всех DSP)
	if err = beforeAuction.RunChain(s.exchangeRepository, dto); err != nil {
		return res, err
	}

	// Получаем список всех подходящих DSP
	if dto.DSPs, err = s.getDSPs(dto); err != nil {
		return res, err
	}

	// Проводим аукцион, получаем победившие ставки
	wonBids, err := s.auction(dto)
	if err != nil {
		return res, err
	}

	// Сохраняем данные для дальнейшей работы и биллинга
	if err = s.saveWonBids(dto, wonBids); err != nil {
		return res, err
	}

	// Если запрос является запросом на барабан, то сохраняем adm в кэш
	if dto.IsClickunder {
		if err = s.saveADMs(ctx, wonBids); err != nil {
			return res, err
		}
	}

	// Подготавливаем ответ в SSP
	bidResponse, err := beforeResponseToSSP.RunChain(dto, wonBids, s.fraudScoreService, s)
	if err != nil {
		return res, err
	}

	// Сохраняем объект аналитики для использования на биллинге
	if err = s.exchangeRepository.SaveAnalyticDTOs(ctx, maps.Values(dto.AnalyticDTOByBid)); err != nil {
		return res, err
	}

	// Сохраняем наш ответ в кафку, чтобы потом проанализировать
	if err = s.eventService.CreateExchangeBidResponseToSSPEvent(ctx, bidResponse); err != nil {
		log.Error(ctx, err)
	}

	// Проверяем, не опоздали ли мы с ответом
	if err = ctx.Err(); err != nil {
		return res, errors.Timeout.Wrap(err)
	}

	return bidResponse, nil
}
