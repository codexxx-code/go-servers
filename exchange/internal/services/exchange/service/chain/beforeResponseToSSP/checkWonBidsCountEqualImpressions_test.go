package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_checkWonBidsCountEqualImpressions_Apply(t *testing.T) {

	tests := []struct {
		name    string
		req     beforeResponseToSSP
		wantErr error
	}{
		{
			name: "1. Одинаковое количество импрешенов и победивших бидов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{

					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{},
							{},
						},
					},
				},
				wonBids: []model.WonBid{
					{},
					{},
				},
			},
			wantErr: nil,
		},
		{
			name: "2. Разное количество импрешенов и победивших бидов",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{

					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{},
							{},
						},
					},
				},
				wonBids: []model.WonBid{
					{},
				},
			},
			wantErr: errors.InternalServer.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &checkWonBidsCountEqualImpressions{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
		})
	}
}
