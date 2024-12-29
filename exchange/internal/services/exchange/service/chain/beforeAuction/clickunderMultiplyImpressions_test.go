package beforeAuction

import (
	"testing"

	"exchange/internal/services/exchange/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/openrtb"
	"pkg/pointer"
	"pkg/testUtils"
)

func Test_setCountOutdoingRequests_Apply(t *testing.T) {

	tests := []struct {
		name       string
		req        beforeAuction
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Баннерный запрос",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								Banner: &openrtb.Banner{},
							},
							{
								Banner: &openrtb.Banner{},
							},
						},
					},
					SSP: sspModel.SSP{
						ClickunderDrumSize: pointer.Pointer(int32(2)),
					},
					IsClickunder: false,
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						Banner: &openrtb.Banner{},
					},
					{
						Banner: &openrtb.Banner{},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "2. Кликандер запрос и размер барабана = 5",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								ID: "1",
							},
						},
					},
					SSP: sspModel.SSP{
						ClickunderDrumSize: pointer.Pointer(int32(5)),
					},
					IsClickunder: true,
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						ID: "1",
					},
					{
						ID: "1",
					},
					{
						ID: "1",
					},
					{
						ID: "1",
					},
					{
						ID: "1",
					},
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := new(clickunderMultiplyImpressions)

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
