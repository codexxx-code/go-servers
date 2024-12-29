package beforeAuction

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_setRequestCurrency_Apply(t *testing.T) {

	tests := []struct {
		name            string
		req             beforeAuction
		requestCurrency string
		wantErr         error
	}{
		{
			name: "1. Несколько импршенов с одинаковой валютой",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								BidFloorCurrency: "USD",
							},
							{
								BidFloorCurrency: "USD",
							},
							{
								BidFloorCurrency: "USD",
							},
						},
					},
					RequestCurrency: "",
				},
			},
			requestCurrency: "USD",
			wantErr:         nil,
		},
		{
			name: "2. Несколько импршенов с разными валютами",
			req: beforeAuction{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Impressions: []openrtb.Impression{
							{
								BidFloorCurrency: "USD",
							},
							{
								BidFloorCurrency: "RUB",
							},
							{
								BidFloorCurrency: "USD",
							},
						},
					},
					RequestCurrency: "",
				},
			},
			requestCurrency: "",
			wantErr:         errors.BadRequest.New(""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &setRequestCurrency{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			if !testUtils.CheckError(t, gotErr, tt.wantErr, false) {
				if tt.req.RequestCurrency != tt.requestCurrency {
					t.Errorf("got %v, want %v", tt.req.RequestCurrency, tt.requestCurrency)
				}
			}
		})
	}
}
