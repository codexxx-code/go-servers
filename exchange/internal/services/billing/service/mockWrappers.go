package service

import (
	"github.com/stretchr/testify/mock"

	billingRepository "exchange/internal/services/billing/repository"
	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
)

//// Обертки над моками для типизации

func (m *MockBillingRepository) MockCreateDSPWinEvent(req billingRepository.CreateDSPEventWinReq,
	err error,
) {
	m.On("CreateDSPWinEvent", mock.Anything, req).
		Return(err)
}

func (m *MockBillingRepository) MockCreateSSPWinEvent(req billingRepository.CreateSSPEventWinReq,
	err error,
) {
	m.On("CreateSSPWinEvent", mock.Anything, req).
		Return(err)
}

func (m *MockExchangeRepository) MockGetDSPResponse(req model.GetDSPResponsesReq,
	dspResponse model.DSPResponse, err error,
) {
	m.On("GetDSPResponses", mock.Anything, req).
		Return([]model.DSPResponse{dspResponse}, err)
}

func (m *MockExchangeRepository) MockGetDSPResponses(req model.GetDSPResponsesReq,
	dspResponses []model.DSPResponse, err error,
) {
	m.On("GetDSPResponses", mock.Anything, req).
		Return(dspResponses, err)
}

func (m *MockDSPService) MockGetDSPs(filters dspModel.GetDSPsReq,
	dsps []dspModel.DSP, err error,
) {
	m.On("GetDSPs", mock.Anything, filters).
		Return(dsps, err)
}

func (m *MockBillingNetwork) MockBillDSP(url string,
	statusCode int, err error,
) {
	m.On("BillDSP", mock.Anything, url).
		Return(statusCode, err)
}
