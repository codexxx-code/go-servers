package utils

import (
	"testing"

	dspModel "exchange/internal/services/dsp/model"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func TestExchangeService_validateResponseFromDSP(t *testing.T) {

	tests := []struct {
		name        string
		bidResponse *openrtb.BidResponse
		dsp         dspModel.DSP
		wantErr     error
	}{
		{
			name:        "1. Ответ без BidResponse",
			bidResponse: nil,
			wantErr:     errors.BadRequest.New(""),
		},
		{
			name: "3. Валюта ответа не совпадает с валютой DSP",
			bidResponse: &openrtb.BidResponse{
				Currency: "USD",
			},
			dsp: dspModel.DSP{
				Currency: "RUB",
			},
			wantErr: errors.BadRequest.New(""),
		},
		{
			name:        "4. BidResponse не прошел валидацию",
			bidResponse: &openrtb.BidResponse{},
			wantErr:     errors.BadRequest.New(""),
		},
		{
			name: "5. Все проверки пройдены",
			bidResponse: &openrtb.BidResponse{
				ID:       "123",
				Currency: "USD",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								ID:    "123",
								ImpID: "123",
							},
						},
					},
				},
			},
			dsp: dspModel.DSP{
				Currency: "USD",
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := ValidateResponseFromDSP(tt.bidResponse, tt.dsp)
			testUtils.CheckError(t, tt.wantErr, gotErr, false)
		})
	}
}
