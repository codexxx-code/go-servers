package http

import (
	"github.com/stretchr/testify/mock"

	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
)

//// Обертки над моками для типизации

func (m *MockExchangeService) MockBidSSP(bidReqFromSSP model.SSPBidReq,
	bidResponseToSSP openrtb.BidResponse, err error,
) {
	m.On("BidSSP", mock.Anything, bidReqFromSSP).
		Return(bidResponseToSSP, err)
}
