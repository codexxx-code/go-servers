package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_addNURLForClickunder_Apply(t *testing.T) {

	generalSettings := settingsModel.Settings{
		Host: "https://test.com",
	}

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Добавление нашего NURL",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					ExchangeID:   "requestID",
					Settings:     generalSettings,
					SSP:          sspModel.SSP{Slug: "ssp_slug"},
					IsClickunder: true,
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
			},
			bidResponse: openrtb.BidResponse{
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								NoticeURL: "https://test.com/billing/requestID?id_type=request&price=${AUCTION_PRICE}&ssp_slug=ssp_slug&url_type=nurl",
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

			r := &addNURLForClickunder{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
