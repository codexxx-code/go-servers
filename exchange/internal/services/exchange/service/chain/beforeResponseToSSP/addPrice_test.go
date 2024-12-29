package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_addPrice_Apply(t *testing.T) {

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Конвертация и проставление цены из выигранных ставок",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates:   currencyRates,
					RequestCurrency: "RUB",
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{},
							},
						},
						{
							Bids: []openrtb.Bid{
								{},
							},
						},
					},
				},
				wonBids: []model.WonBid{
					{
						BidPointer: model.BidPointer{
							BidResponse: openrtb.BidResponse{
								Currency: "RUB",
							},
						},
						BillingPriceInDSPCurrency: decimal.NewFromFloat(123.45),
					},
					{
						BidPointer: model.BidPointer{
							BidResponse: openrtb.BidResponse{
								Currency: "EUR",
							},
						},
						BillingPriceInDSPCurrency: decimal.NewFromFloat(1),
					},
				},
			},
			bidResponse: openrtb.BidResponse{
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								Price: decimal.NewFromFloat(123.45),
							},
						},
					},
					{
						Bids: []openrtb.Bid{
							{
								Price: decimal.NewFromFloat(110),
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

			r := &addPrice{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
