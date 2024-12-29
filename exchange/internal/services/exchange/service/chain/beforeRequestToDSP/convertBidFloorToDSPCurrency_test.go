package beforeRequestToDSP

import (
	"testing"

	dspModel "exchange/internal/services/dsp/model"
	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_convertBidFloorToDSPCurrency_Apply(t *testing.T) {
	generalCurrencyRates := map[string]decimal.Decimal{
		"USD": decimal.NewFromFloat(1),
		"RUB": decimal.NewFromFloat(0.01),
		"EUR": decimal.NewFromFloat(1.1),
	}

	tests := []struct {
		name       string
		req        beforeRequestToDSP
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Валюта запроса не совпадает с валютой DSP, два импрешена",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				BidRequest: openrtb.BidRequest{
					Currencies: []string{"USD", "EUR", "THB", "RUB"},
					Impressions: []openrtb.Impression{
						{
							BidFloor:         decimal.NewFromFloat(1.5),
							BidFloorCurrency: "USD",
						},
						{
							BidFloor:         decimal.NewFromFloat(2),
							BidFloorCurrency: "EUR",
						},
					},
				},
				dsp: dspModel.DSP{
					Currency: "RUB",
				},
			},
			bidRequest: openrtb.BidRequest{
				Currencies: []string{"RUB"},
				Impressions: []openrtb.Impression{
					{
						BidFloor:         decimal.NewFromInt(150),
						BidFloorCurrency: "RUB",
					},
					{
						BidFloor:         decimal.NewFromInt(220),
						BidFloorCurrency: "RUB",
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "2. Валюта запроса совпадает с валютой DSP, два импрешена",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				BidRequest: openrtb.BidRequest{
					Currencies: []string{"USD", "EUR", "THB", "RUB"},
					Impressions: []openrtb.Impression{
						{
							BidFloor:         decimal.NewFromFloat(150),
							BidFloorCurrency: "RUB",
						},
						{
							BidFloor:         decimal.NewFromFloat(220),
							BidFloorCurrency: "RUB",
						},
					},
				},
				dsp: dspModel.DSP{
					Currency: "RUB",
				},
			},
			bidRequest: openrtb.BidRequest{
				Currencies: []string{"RUB"},
				Impressions: []openrtb.Impression{
					{
						BidFloor:         decimal.NewFromInt(150),
						BidFloorCurrency: "RUB",
					},
					{
						BidFloor:         decimal.NewFromInt(220),
						BidFloorCurrency: "RUB",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &convertBidFloorToDSPCurrency{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
