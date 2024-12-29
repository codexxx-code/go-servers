package beforeAnalytic

import (
	"testing"

	"exchange/internal/services/exchange/model"
	"pkg/errors"
	"pkg/openrtb"
	"pkg/testUtils"
)

func Test_logGeoByIP_Apply(t *testing.T) {

	tests := []struct {
		name        string
		req         beforeAnalytic
		bidRequest  openrtb.BidRequest
		wantErr     error
		mockActions func(mocks)
	}{
		{
			name: "1. Пустой объект device",
			req: beforeAnalytic{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Device: nil,
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Device: nil,
			},
			wantErr:     errors.BadRequest.New(""),
			mockActions: nil,
		},
		{
			name: "2. Пустой IP",
			req: beforeAnalytic{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Device: &openrtb.Device{
							IP: "",
						},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Device: &openrtb.Device{
					IP: "",
				},
			},
			wantErr:     errors.BadRequest.New(""),
			mockActions: nil,
		},
		{
			name: "4. Пустой User",
			req: beforeAnalytic{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Device: &openrtb.Device{
							IP: "11",
						},
						User: nil,
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Device: &openrtb.Device{
					IP: "11",
				},
			},
			wantErr:     errors.BadRequest.New(""),
			mockActions: nil,
		},
		{
			name: "4. Ошибка при получении страны по IP",
			req: beforeAnalytic{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Device: &openrtb.Device{
							IP: "1",
						},
						User: &openrtb.User{},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Device: &openrtb.Device{
					IP: "1",
				},
				User: &openrtb.User{},
			},
			wantErr: errors.InternalServer.New(""),
			mockActions: func(m mocks) {
				m.exchangeRepository.MockGetCountryByIP("1",
					"", errors.InternalServer.New(""))
			},
		},
		{
			name: "5. Успешное выполнение",
			req: beforeAnalytic{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Device: &openrtb.Device{
							IP: "1",
						},
						User: &openrtb.User{},
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Device: &openrtb.Device{
					IP: "1",
					Geo: &openrtb.Geo{
						Country: "RUS",
					},
				},
				User: &openrtb.User{
					Geo: &openrtb.Geo{
						Country: "RUS",
					},
				},
			},
			wantErr: nil,
			mockActions: func(m mocks) {
				m.exchangeRepository.MockGetCountryByIP("1",
					"RUS", nil)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			mocks := getMocks(t)

			r := &checkGeoByIP{
				exchangeRepository: mocks.exchangeRepository,
			}

			// Выполняем действия с моками
			if tt.mockActions != nil {
				tt.mockActions(mocks)
			}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
