package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

var currencyRates = map[string]decimal.Decimal{
	"RUB": decimal.NewFromFloat(0.01),
	"EUR": decimal.NewFromFloat(1.1),
}

func Test_addPriceForClickunder_Apply(t *testing.T) {

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Проставление цены из BillingPriceInDSPCurrency в NURL для кликандер SSP",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates:   currencyRates,
					RequestCurrency: "RUB",
					IsClickunder:    true,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
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
								Price: decimal.NewFromFloat(233.45),
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

			r := &addPriceForClickunder{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
