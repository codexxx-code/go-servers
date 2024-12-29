package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_add_ADM_Apply(t *testing.T) {

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Добавление ADM в ответ на баннерный запрос",
			req: beforeResponseToSSP{
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
								SeatBids: []openrtb.SeatBid{
									{
										Bids: []openrtb.Bid{
											{
												AdMarkup: "AdMarkupFromDSP",
											},
										},
									},
								},
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
								AdMarkup: "AdMarkupFromDSP",
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

			r := &addADM{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
