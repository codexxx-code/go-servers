package beforeRequestToDSP

import (
	"testing"

	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_applyAuctionType_Apply(t *testing.T) {

	tests := []struct {
		name       string
		req        beforeRequestToDSP
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Настройка DSP = аукцион первой цены и значение AuctionType = 2",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						AuctionType: 2,
					},
				},
				dsp: dspModel.DSP{
					AuctionSecondPrice: false,
				},
			},
			bidRequest: openrtb.BidRequest{
				AuctionType: 1,
			},
			wantErr: nil,
		},
		{
			name: "2. Настройка DSP = аукцион второй цены и значение AuctionType = 1",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						AuctionType: 1,
					},
				},
				dsp: dspModel.DSP{
					AuctionSecondPrice: true,
				},
			},
			bidRequest: openrtb.BidRequest{
				AuctionType: 2,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &applyAuctionType{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
