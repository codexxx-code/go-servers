package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
	"pkg/testUtils"
	"pkg/uuid"
)

func Test_initBids_Apply(t *testing.T) {

	tests := []struct {
		name         string
		req          beforeResponseToSSP
		bidResponse  openrtb.BidResponse
		mockedValues []string
		wantErr      error
	}{
		{
			name: "1. Добавление двух бидов на два импрешена",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{

					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								ID: "imp1",
							},
							{
								ID: "imp2",
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
								ID:    "1",
								ImpID: "imp1",
							},
						},
					},
					{
						Bids: []openrtb.Bid{
							{
								ID:    "2",
								ImpID: "imp2",
							},
						},
					},
				},
			},
			mockedValues: []string{"1", "2"},
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &initBids{}

			uuid.AddMockValues(tt.mockedValues...)

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
