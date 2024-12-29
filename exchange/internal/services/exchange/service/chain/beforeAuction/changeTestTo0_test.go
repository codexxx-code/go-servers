package beforeAuction

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_changeTestTo0_Apply(t *testing.T) {

	tests := []struct {
		name       string
		req        beforeAuction
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Test = 1",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Test: 1,
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Test: 0,
			},
			wantErr: nil,
		},
		{
			name: "2. Test = 0",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Test: 0,
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Test: 0,
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &changeTestTo0{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
