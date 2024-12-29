package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_checkMaxPrice_Apply(t *testing.T) {

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
			name: "1. Ответ DSP с ценой 1000 рублей за один из бидов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(100),
								},
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
			name: "2. Ответ DSP с ценой 5 евро за один из бидов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(5),
								},
								{
									Price: decimal.NewFromFloat(0.3),
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
			name: "3. Ответ DSP с ценой 500 рублей за один из бидов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(500),
								},
								{
									Price: decimal.NewFromInt(499),
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
			name: "4. Ответ DSP с ценами до 500 рублей",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(234),
								},
								{
									Price: decimal.NewFromInt(499),
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
			name: "5. Ответ DSP с ценой до 5 долларов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					CurrencyRates: generalCurrencyRates,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price: decimal.NewFromInt(1),
								},
								{
									Price: decimal.NewFromInt(2),
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

			r := &checkMaxPrice{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, true)
		})
	}
}
