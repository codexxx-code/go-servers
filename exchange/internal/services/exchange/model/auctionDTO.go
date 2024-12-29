package model

import (
	"context"
	"sync"
	"time"

	"exchange/internal/enum/sourceTrafficType"
	dspModel "exchange/internal/services/dsp/model"
	settingsModel "exchange/internal/services/setting/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/decimal"
	"pkg/openrtb"
)

type AuctionDTO struct {
	Ctx context.Context // Глобальный контекст запроса

	OriginalBidRequest openrtb.BidRequest // Изначальный неизмененный BidRequest. ЕГО МЫ НЕ МЕНЯЕМ
	BidRequest         openrtb.BidRequest // Рабочий BidRequest, который может меняться

	SSP           sspModel.SSP               // SSP, которая инициализировала запрос
	Settings      settingsModel.Settings     // Настройки системы
	DSPs          []dspModel.DSP             // Все DSP, подходящие для этого запроса
	CurrencyRates map[string]decimal.Decimal // Курсы валют

	ExchangeID                   string                     // Наш идентификатор запроса
	MappingExchangeImpressionIDs map[string]string          // Маппинг наших идентификаторов импрешенов на идентификаторы SSP
	MappingImpressionsByDSPs     MappingImpressionsByDSPsMu // Маппинг идентификаторов импрешенов на идентификаторы. Идентификатор DSP - рандомный идентификатор импрешена - наш идентификатор импрешена

	IsClickunder      bool                                // Является ли запрос кликандером
	DrumSize          *int32                              // Размер барабана
	RequestCurrency   string                              // Валюта запроса
	SourceTrafficType sourceTrafficType.SourceTrafficType // Тип запроса по источнику
	RequestTimeout    time.Duration                       // Время, которое есть у программы на запрос
	GeoCountry        string                              // Страна запроса
	PublisherID       string                              // Идентификатор паблишера

	AnalyticDTOByImpression             map[string]AnalyticDTO                // DTO для аналитики в разбивке по импрешенам
	AnalyticDTOByDSPRequestByImpression AnalyticDTOByDSPRequestByImpressionMu // DTO для аналитики в разбивке по запросам в DSP и импрешенам
	AnalyticDTOByBid                    map[string]AnalyticDTO                // DTO для аналитики в разбивке по бидам
}

type AnalyticDTOByDSPRequestByImpressionMu struct {
	*sync.RWMutex
	AnalyticDTOByDSPRequestByImpression map[string]map[string]AnalyticDTO
}

func NewAnalyticDTOByDSPRequestByImpressionMu() AnalyticDTOByDSPRequestByImpressionMu {
	return AnalyticDTOByDSPRequestByImpressionMu{
		RWMutex:                             &sync.RWMutex{},
		AnalyticDTOByDSPRequestByImpression: make(map[string]map[string]AnalyticDTO),
	}
}

func (m *AnalyticDTOByDSPRequestByImpressionMu) Get(dspSlug, impressionID string) (dto AnalyticDTO, ok bool) {
	m.RLock()
	defer m.RUnlock()
	dto, ok = m.AnalyticDTOByDSPRequestByImpression[dspSlug][impressionID]
	return dto, ok
}

func (m *AnalyticDTOByDSPRequestByImpressionMu) Set(dspSlug, impressionID string, dto AnalyticDTO) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.AnalyticDTOByDSPRequestByImpression[dspSlug]; !ok {
		m.AnalyticDTOByDSPRequestByImpression[dspSlug] = make(map[string]AnalyticDTO)
	}
	m.AnalyticDTOByDSPRequestByImpression[dspSlug][impressionID] = dto
}

type MappingImpressionsByDSPsMu struct {
	*sync.RWMutex
	Mapping map[string]map[string]string
}

func NewMappingImpressionsByDSPsMu() MappingImpressionsByDSPsMu {
	return MappingImpressionsByDSPsMu{
		RWMutex: &sync.RWMutex{},
		Mapping: make(map[string]map[string]string),
	}
}

func (m *MappingImpressionsByDSPsMu) Get(dspSlug, randomImpressionID string) (exchangeImpressionID string, ok bool) {
	m.RLock()
	defer m.RUnlock()
	exchangeImpressionID, ok = m.Mapping[dspSlug][randomImpressionID]
	return exchangeImpressionID, ok
}

func (m *MappingImpressionsByDSPsMu) Set(dspSlug, impressionID, exchangeImpressionID string) {
	m.Lock()
	defer m.Unlock()
	if _, ok := m.Mapping[dspSlug]; !ok {
		m.Mapping[dspSlug] = make(map[string]string)
	}
	m.Mapping[dspSlug][impressionID] = exchangeImpressionID
}
