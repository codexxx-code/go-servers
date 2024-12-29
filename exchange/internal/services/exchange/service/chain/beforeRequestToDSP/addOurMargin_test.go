package beforeRequestToDSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_addOurMargin_Apply(t *testing.T) {
	tests := []struct {
		name       string
		req        beforeRequestToDSP
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Настройка маржи = 0",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					Settings: settingsModel.Settings{
						Margin: decimal.NewFromInt(0),
					},
				},
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor: decimal.NewFromInt(10),
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor: decimal.NewFromInt(10),
					},
				},
			},
		},
		{
			name: "2. Настройка маржи = 0.5",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					Settings: settingsModel.Settings{
						Margin: decimal.NewFromFloat(0.5),
					},
				},
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor: decimal.NewFromInt(10),
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor: decimal.NewFromInt(15),
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := new(addOurMargin)

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
