package service

import (
	"testing"

	"github.com/stretchr/testify/mock"

	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/decimal"
	"pkg/openrtb"
)

// Обертки над моками для типизации

type Mocks struct {
	ExchangeRepository *MockExchangeRepository
	ExchangeNetwork    *MockExchangeNetwork
	SSPService         *MockSSPService
	DSPService         *MockDSPService
	CurrencyService    *MockCurrencyService
	SettingService     *MockSettingService
	EventService       *MockEventService
}

func NewMocks(t *testing.T) Mocks {
	return Mocks{
		ExchangeRepository: NewMockExchangeRepository(t),
		ExchangeNetwork:    NewMockExchangeNetwork(t),
		SSPService:         NewMockSSPService(t),
		DSPService:         NewMockDSPService(t),
		CurrencyService:    NewMockCurrencyService(t),
		SettingService:     NewMockSettingService(t),
		EventService:       NewMockEventService(t),
	}
}

func (m *MockExchangeRepository) MockCreateDSPResponse(i int, req model.DSPResponse,
	err error,
) {
	m.On("CreateDSPResponse", mock.Anything, i, req).Return(err)
}

func (m *MockExchangeRepository) MockGetADM(id string,
	res string, err error,
) {
	m.On("GetADM", mock.Anything, id).Return(res, err)
}

func (m *MockExchangeRepository) MockSaveADM(id, adm string,
	err error,
) {
	m.On("SaveADM", mock.Anything, id, adm).Return(err)
}

func (m *MockExchangeRepository) MockGetCountryByIP(req string,
	res string, err error,
) {
	m.On("GetCountryByIP", req).Return(res, err)
}

func (m *MockExchangeRepository) MockGetPublisherVisibility(req string,
	res model.PublisherVisibility, err error,
) {
	m.On("GetPublisherVisibility", mock.Anything, req).Return(res, err)
}

func (m *MockExchangeNetwork) MockSendBidRequestToDSP(requestNumber int, url string, req openrtb.BidRequest,
	res openrtb.BidResponse, code int, err error,
) {
	m.On("SendBidRequestToDSP", mock.Anything, requestNumber, url, req).Return(res, code, err)
}

func (m *MockSSPService) MockGetSSP(req sspModel.GetSSPsReq,
	res sspModel.SSP, err error,
) {
	m.On("GetSSPs", mock.Anything, req).Return([]sspModel.SSP{res}, err)
}

func (m *MockDSPService) MockGetDSPs(req dspModel.GetDSPsReq,
	res []dspModel.DSP, err error,
) {
	m.On("GetDSPs", mock.Anything, req).Return(res, err)
}

func (m *MockCurrencyService) MockGetRates(
	res map[string]decimal.Decimal, err error,
) {
	m.On("GetRates", mock.Anything).Return(res, err)
}

func (m *MockSettingService) MockGetSettings(
	res settingsModel.Settings, err error,
) {
	m.On("GetSettings", mock.Anything).Return(res, err)
}

func (m *MockEventService) MockCreateSSPBidRequestEvent(req openrtb.BidRequest,
	err error,
) {
	m.On("CreateSSPBidRequestToExchangeEvent", mock.Anything, req).Return(err)
}

func (m *MockEventService) MockCreateExchangeBidResponseEvent(req openrtb.BidResponse,
	err error,
) {
	m.On("CreateExchangeBidResponseToSSPEvent", mock.Anything, req).Return(err)
}

func (m *MockEventService) MockCreateExchangeBidRequestToDSPEvent(req openrtb.BidRequest, i int,
	err error,
) {
	m.On("CreateExchangeBidRequestToDSPEvent", mock.Anything, i, req).Return(err)
}
