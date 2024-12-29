package http

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"exchange/internal/services/billing/model"
	"exchange/internal/services/billing/service"
)

//// Обертки над моками для типизации

type mocks struct {
	exchangeRepository    *service.MockExchangeRepository
	dspService            *service.MockDSPService
	billingRepository     *service.MockBillingRepository
	billingNetwork        *service.MockBillingNetwork
	analyticWriterSerivce *service.MockAnalyticWriterService
}

func getMocks(t *testing.T) mocks {
	return mocks{
		exchangeRepository:    service.NewMockExchangeRepository(t),
		dspService:            service.NewMockDSPService(t),
		billingRepository:     service.NewMockBillingRepository(t),
		billingNetwork:        service.NewMockBillingNetwork(t),
		analyticWriterSerivce: service.NewMockAnalyticWriterService(t),
	}
}

func (m *MockBillingService) MockBillURL(billUrlReq model.BillURLReq,
	err error,
) {
	m.On("BillURL", mock.Anything, billUrlReq).
		Return(err)
}
