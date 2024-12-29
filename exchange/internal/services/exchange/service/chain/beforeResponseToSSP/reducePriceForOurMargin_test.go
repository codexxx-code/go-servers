package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_reducePriceForOurMargin_Apply(t *testing.T) {

	settings := settingsModel.Settings{
		Margin: decimal.NewFromFloat(0.1),
	}

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Снижение ставки ответа в SSP на 10% маржи",
			req: beforeResponseToSSP{
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromFloat(100),
								},
							},
						},
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromFloat(1),
								},
							},
						},
					},
				},
				AuctionDTO: &model.AuctionDTO{
					Settings: settings,
				},
				chainSettings: chainSettings{
					priceAlreadySet: true,
				},
			},
			bidResponse: openrtb.BidResponse{
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								Price: decimal.NewFromFloat(90),
							},
						},
					},
					{
						Bids: []openrtb.Bid{
							{
								Price: decimal.NewFromFloat(0.9),
							},
						},
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &reducePriceForOurMargin{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
