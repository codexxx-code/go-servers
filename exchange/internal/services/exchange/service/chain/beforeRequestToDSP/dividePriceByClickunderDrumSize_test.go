package beforeRequestToDSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/pointer"
	"pkg/testUtils"
)

func Test_dividePriceByMultiplicationFactor_Apply(t *testing.T) {
	tests := []struct {
		name       string
		req        beforeRequestToDSP
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Несколько импрешенов и кликандер с размером барабана = 5",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					IsClickunder: true,
					SSP: sspModel.SSP{
						ClickunderDrumSize: pointer.Pointer(int32(5)),
					},
				},
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor: decimal.NewFromFloat(5),
						},
						{
							BidFloor: decimal.NewFromFloat(5),
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor: decimal.NewFromInt(1),
					},
					{
						BidFloor: decimal.NewFromInt(1),
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "2. Несколько импрешенов и не кликандер",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					IsClickunder: false,
					SSP: sspModel.SSP{
						ClickunderDrumSize: pointer.Pointer(int32(5)),
					},
				},
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor: decimal.NewFromFloat(5),
						},
						{
							BidFloor: decimal.NewFromFloat(5),
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor: decimal.NewFromFloat(5),
					},
					{
						BidFloor: decimal.NewFromFloat(5),
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &dividePriceByClickunderDrumSize{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
