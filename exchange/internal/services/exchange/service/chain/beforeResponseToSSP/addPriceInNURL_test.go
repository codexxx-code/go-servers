package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_addPriceInNURL_Apply(t *testing.T) {

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Конвертация, суммирование и проставление цены из выигранных ставок",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					IsClickunder: true,
				},
				bidResponse: openrtb.BidResponse{
					SeatBids: []openrtb.SeatBid{
						{
							Bids: []openrtb.Bid{
								{
									Price:     decimal.NewFromFloat(123.45),
									NoticeURL: "https://host.com/nurl?price=${AUCTION_PRICE}",
								},
							},
						},
					},
				},
				chainSettings: chainSettings{
					nurlAlreadySet: true,
				},
			},
			bidResponse: openrtb.BidResponse{
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								Price:     decimal.NewFromFloat(123.45),
								NoticeURL: "https://host.com/nurl?price=${AUCTION_PRICE}&bid_price=123.45",
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

			r := &addPriceInNURL{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
