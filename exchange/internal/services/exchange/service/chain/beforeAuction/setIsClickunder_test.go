package beforeAuction

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_setIsClickunder_Apply(t *testing.T) {

	tests := []struct {
		name         string
		req          beforeAuction
		isClickunder bool
		wantErr      error
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
						},
					},
					IsClickunder: false,
				},
			},
			isClickunder: false,
			wantErr:      nil,
		},
		{
			name: "2. Видео запрос",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								Video: &openrtb.Video{},
							},
						},
					},
					IsClickunder: false,
				},
			},
			isClickunder: false,
			wantErr:      nil,
		},
		{
			name: "3. Аудио запрос",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								Audio: &openrtb.Audio{},
							},
						},
					},
					IsClickunder: false,
				},
			},
			isClickunder: false,
			wantErr:      nil,
		},
		{
			name: "4. Нативный запрос",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								Native: &openrtb.Native{},
							},
						},
					},
					IsClickunder: false,
				},
			},
			isClickunder: false,
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &setIsClickunder{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			if tt.req.IsClickunder != tt.isClickunder {
				t.Errorf("setIsClickunder.Apply() = %v, want %v", tt.req.IsClickunder, tt.isClickunder)
			}
		})
	}
}
