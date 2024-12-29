package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_checkMaxPriceClickunder_Apply(t *testing.T) {

	generalCurrencyRates := map[string]decimal.Decimal{
		"USD": decimal.NewFromFloat(1),
		"RUB": decimal.NewFromFloat(0.01),
		"EUR": decimal.NewFromFloat(1.1),
	}

	tests := []struct {
		name    string
		req     beforeResponseToSSP
		wantErr error
	}{
		{
			name: "1. Ответ DSP с ценой 1100 рублей за бид",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
					IsClickunder:  true,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(1000),
								},
							},
						},
					},
					Currency: "RUB",
				},
			},
			wantErr: errors.BadRequest.Wrap(ErrPriceIsTooHigh),
		},
		{
			name: "2. Ответ DSP с ценой 12 евро за бид",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
					IsClickunder:  true,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(12),
								},
							},
						},
					},
					Currency: "EUR",
				},
			},
			wantErr: errors.BadRequest.Wrap(ErrPriceIsTooHigh),
		},
		{
			name: "3. Ответ DSP с ценой до 500 рублей",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
					IsClickunder:  true,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(234),
								},
							},
						},
					},
					Currency: "RUB",
				},
			},
			wantErr: nil,
		},
		{
			name: "4. Ответ DSP с ценой до 5 долларов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
					IsClickunder:  true,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(1),
								},
							},
						},
					},
					Currency: "USD",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &checkMaxPriceClickunder{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, true)
		})
	}
}
