package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
	"pkg/testUtils"
	"pkg/uuid"
)

func Test_initResponse_Apply(t *testing.T) {
	tests := []struct {
		name         string
		req          beforeResponseToSSP
		bidResponse  openrtb.BidResponse
		mockedValues []string
		wantErr      error
	}{
		{
			name: "1. Успешное создание идентификатора запроса в bidResponse от DSP и создание bidderID",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						ID: "123",
					},
				},
				bidResponse: openrtb.BidResponse{
					ID: "456",
				},
			},
			bidResponse: openrtb.BidResponse{
				ID:    "123",
				BidID: "1",
			},
			mockedValues: []string{"1"},
			wantErr:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &initResponse{}

			uuid.AddMockValues(tt.mockedValues...)

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
