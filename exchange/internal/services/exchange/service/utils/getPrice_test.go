package utils

import (
	"testing"

	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func TestExchangeService_getPrice(t *testing.T) {

	defaultSettings := settingsModel.Settings{
		EmptySecondPriceReduceCoef: decimal.NewFromFloat(0.1),
	}

	currencyRates := map[string]decimal.Decimal{
		"USD": decimal.NewFromFloat(1),
		"RUB": decimal.NewFromFloat(0.01),
		"EUR": decimal.NewFromFloat(1.1),
	}

	test1Response1 := openrtb.BidResponse{
		SeatBids: []openrtb.SeatBid{
			{
				Bids: []openrtb.Bid{
					{
						ImpID: "1",
						Price: decimal.NewFromFloat(10), // 10 рублей - 1 импрешен - 2
					},
				},
			},
			{
				Bids: []openrtb.Bid{
					{
						ImpID: "2",
						Price: decimal.NewFromFloat(20), // 20 рублей - 2 импрешен - 2
					},
					{
						ImpID: "3",
						Price: decimal.NewFromFloat(33), // 33 рубля - 3 импрешен - 1
					},
				},
			},
		},
		Currency: "RUB",
	}
	test1Response2 := openrtb.BidResponse{
		SeatBids: []openrtb.SeatBid{
			{
				Bids: []openrtb.Bid{
					{
						ImpID: "1",
						Price: decimal.NewFromFloat(0.44), // 0.44 доллара = 44 рубля - 1 импрешен - 1
					},
				},
			},
		},
		Currency: "USD",
	}
	test1Response3 := openrtb.BidResponse{
		SeatBids: []openrtb.SeatBid{
			{
				Bids: []openrtb.Bid{
					{
						ImpID: "2",
						Price: decimal.NewFromFloat(0.5), // 0.5 евро = 55 рублей - 2 импрешен - 1
					},
				},
			},
		},
		Currency: "EUR",
	}

	type args struct {
		bidPointers []model.BidPointer
		bidRequest  openrtb.BidRequest
	}
	tests := []struct {
		name    string
		args    args
		want    []model.WonBid
		wantErr error
	}{
		{
			name: "1. 3 импрешена в запросе, 5 бидов от DSP",
			args: args{
				bidPointers: []model.BidPointer{
					{
						ExchangeBidID:        "1",
						BidResponse:          test1Response1,
						SeatBidIndex:         0,
						BidIndex:             0,
						DSPSlug:              "dsp1",
						IsAuctionSecondPrice: true,
					},
					{
						ExchangeBidID:        "2",
						BidResponse:          test1Response1,
						SeatBidIndex:         1,
						BidIndex:             0,
						DSPSlug:              "dsp1",
						IsAuctionSecondPrice: true,
					},
					{
						ExchangeBidID:        "3",
						BidResponse:          test1Response1,
						SeatBidIndex:         1,
						BidIndex:             1,
						DSPSlug:              "dsp1",
						IsAuctionSecondPrice: true,
					},
					{
						ExchangeBidID:        "4",
						BidResponse:          test1Response2,
						SeatBidIndex:         0,
						BidIndex:             0,
						DSPSlug:              "dsp2",
						IsAuctionSecondPrice: false,
					},
					{
						ExchangeBidID:        "5",
						BidResponse:          test1Response3,
						SeatBidIndex:         0,
						BidIndex:             0,
						DSPSlug:              "dsp3",
						IsAuctionSecondPrice: false,
					},
				},
				bidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							ID: "1",
						},
						{
							ID: "2",
						},
						{
							ID: "3",
						},
					},
				},
			},
			want: []model.WonBid{
				{
					BidPointer: model.BidPointer{
						ExchangeBidID:        "4",
						BidResponse:          test1Response2,
						SeatBidIndex:         0,
						BidIndex:             0,
						DSPSlug:              "dsp2",
						IsAuctionSecondPrice: false,
					},
					BillingPriceInDSPCurrency: decimal.NewFromFloat(0.44), // Своя ставка
				},
				{
					BidPointer: model.BidPointer{
						ExchangeBidID:        "5",
						BidResponse:          test1Response3,
						SeatBidIndex:         0,
						BidIndex:             0,
						DSPSlug:              "dsp3",
						IsAuctionSecondPrice: false,
					},
					BillingPriceInDSPCurrency: decimal.NewFromFloat(0.5), // Ставка второго победившего бида = 55 рублей
				},
				{
					BidPointer: model.BidPointer{
						ExchangeBidID:        "3",
						BidResponse:          test1Response1,
						SeatBidIndex:         1,
						BidIndex:             1,
						DSPSlug:              "dsp1",
						IsAuctionSecondPrice: true,
					},
					BillingPriceInDSPCurrency: decimal.NewFromFloat(29.7), // Вторая цена, но из-за отсутствия второго бида, -10%
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetWinBids(tt.args.bidPointers, defaultSettings, currencyRates, tt.args.bidRequest)
			testUtils.CheckError(t, err, tt.wantErr, false)
			testUtils.CheckStruct(t, got, tt.want)
		})
	}
}
