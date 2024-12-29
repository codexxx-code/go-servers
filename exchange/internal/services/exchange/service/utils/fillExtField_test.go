package utils

import (
	"testing"

	"pkg/openrtb"
	"pkg/testUtils"
)

func TestExchangeService_fillExtField(t *testing.T) {

	type args struct {
		req       openrtb.BidRequest
		newValues map[string]any
	}
	tests := []struct {
		name    string
		args    args
		wantRes openrtb.BidRequest
		wantErr error
	}{
		{
			name: "1. Добавление ssp-slug в BidRequest с пустым ext",
			args: args{
				req: openrtb.BidRequest{
					ID: "1",
				},
				newValues: map[string]any{"ssp-slug": "ssp-slug"},
			},
			wantRes: openrtb.BidRequest{
				ID:  "1",
				Ext: []byte(`{"ssp-slug":"ssp-slug"}`),
			},
		},
		{
			name: "2. Добавление ssp-slug в BidRequest с заполненным другими полями ext",
			args: args{
				req: openrtb.BidRequest{
					ID:  "1",
					Ext: []byte(`{"some-field":"some-value"}`),
				},
				newValues: map[string]any{"ssp-slug": "ssp-slug"},
			},
			wantRes: openrtb.BidRequest{
				ID:  "1",
				Ext: []byte(`{"some-field":"some-value","ssp-slug":"ssp-slug"}`),
			},
		},
		{
			name: "3. Добавление ssp-slug в BidRequest с уже заполненным полем ssp-slug",
			args: args{
				req: openrtb.BidRequest{
					ID:  "1",
					Ext: []byte(`{"some-field":"some-value","ssp-slug":"some-value"}`),
				},
				newValues: map[string]any{"ssp-slug": "ssp-slug"},
			},
			wantRes: openrtb.BidRequest{
				ID:  "1",
				Ext: []byte(`{"some-field":"some-value","ssp-slug":"ssp-slug"}`),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			gotRes, err := FillExtField(tt.args.req, tt.args.newValues)
			testUtils.CheckError(t, err, tt.wantErr, false)
			testUtils.CheckStruct(t, gotRes, tt.wantRes)
		})
	}
}
