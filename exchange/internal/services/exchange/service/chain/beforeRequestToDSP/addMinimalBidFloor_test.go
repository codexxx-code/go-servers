package beforeRequestToDSP

import (
	"testing"

	"pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_addMinimalBidFloor_Apply(t *testing.T) {

	tests := []struct {
		name       string
		req        beforeRequestToDSP
		bidRequest openrtb.BidRequest
		wantErr    error
	}{
		{
			name: "1. Пустой bidfloor и валюта DSP = RUB",
			req: beforeRequestToDSP{
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor:         decimal.Zero,
							BidFloorCurrency: "RUB",
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor:         decimal.NewFromInt(10),
						BidFloorCurrency: "RUB",
					},
				},
			},
		},
		{
			name: "2. Пустой bidfloor и неизвестная валюта",
			req: beforeRequestToDSP{
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor:         decimal.Zero,
							BidFloorCurrency: "invalid",
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor:         decimal.Zero,
						BidFloorCurrency: "invalid",
					},
				},
			},
		},
		{
			name: "3. Не пустой bidfloor и валюта DSP = RUB",
			req: beforeRequestToDSP{
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor:         decimal.NewFromInt(20),
							BidFloorCurrency: "RUB",
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor:         decimal.NewFromInt(20),
						BidFloorCurrency: "RUB",
					},
				},
			},
		},
		{
			name: "4. Не пустой bidfloor и неизвестная валюта",
			req: beforeRequestToDSP{
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							BidFloor:         decimal.NewFromInt(20),
							BidFloorCurrency: "invalid",
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						BidFloor:         decimal.NewFromInt(20),
						BidFloorCurrency: "invalid",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := new(addMinimalBidFloor)

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
