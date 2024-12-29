package service

import (
	"testing"

	"github.com/stretchr/testify/mock"

	"pkg/openrtb"
)

type Mocks struct {
	EventRepository    *MockEventRepository
	ExchangeRepository *MockExchangeRepository
}

func NewMocks(t *testing.T) *Mocks {
	return &Mocks{
		EventRepository:    NewMockEventRepository(t),
		ExchangeRepository: NewMockExchangeRepository(t),
	}
}

func (m *MockEventRepository) MockCreateSSPBidRequestEvent(req openrtb.BidRequest,
	err error,
) {
	m.On("CreateSSPBidRequestToExchangeEvent", mock.Anything, req).Return(err)
}

func (m *MockEventRepository) MockCreateExchangeBidResponseEvent(req openrtb.BidResponse,
	err error,
) {
	m.On("CreateExchangeBidResponseToSSPEvent", mock.Anything, req).Return(err)
}
