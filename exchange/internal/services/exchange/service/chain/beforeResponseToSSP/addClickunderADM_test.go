package beforeResponseToSSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	sspModel "exchange/internal/services/ssp/model"
	"pkg/openrtb"
	"pkg/pointer"
	"pkg/testUtils"
)

func Test_add_clickunder_ADM_Apply(t *testing.T) {

	generalSettings := settingsModel.Settings{
		ShowcaseURL: "https://showcase.com",
	}

	tests := []struct {
		name        string
		req         beforeResponseToSSP
		bidResponse openrtb.BidResponse
		wantErr     error
	}{
		{
			name: "1. Добавление XML ADM в ответ на кликандер запрос",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Site: &openrtb.Site{
							Inventory: openrtb.Inventory{
								ID: "siteID",
							},
						},
					},
					IsClickunder: true,
					Settings:     generalSettings,
					SSP: sspModel.SSP{
						ClickunderADMFormat: pointer.Pointer("<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<ad><popunderAd><url><![CDATA[${ADM_URL}]]></url></popunderAd></ad>"),
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
				wonBids: []model.WonBid{
					{

						BidPointer: model.BidPointer{
							ExchangeBidID: "bid1",
						},
					},
					{
						BidPointer: model.BidPointer{
							ExchangeBidID: "bid2",
						},
					},
					{
						BidPointer: model.BidPointer{
							ExchangeBidID: "bid3",
						},
					},
				},
			},
			bidResponse: openrtb.BidResponse{
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								AdMarkup: "<?xml version=\"1.0\" encoding=\"ISO-8859-1\"?>\n<ad><popunderAd><url><![CDATA[https://showcase.com/game/happy-bucket/play?ad_type=banner&adm_id=bid1,bid2,bid3&is_adult=true]]></url></popunderAd></ad>",
							},
						},
					},
				},
			},
			wantErr: nil,
		},
		{
			name: "2. Добавление URL ADM в ответ на кликандер запрос",
			req: beforeResponseToSSP{
				AuctionDTO: &model.AuctionDTO{
					BidRequest: openrtb.BidRequest{
						Site: &openrtb.Site{
							Inventory: openrtb.Inventory{
								ID: "siteID",
							},
						},
					},
					Settings:     generalSettings,
					IsClickunder: true,
					SSP: sspModel.SSP{
						ClickunderADMFormat: pointer.Pointer("${ADM_URL}"),
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
				wonBids: []model.WonBid{
					{
						BidPointer: model.BidPointer{
							ExchangeBidID: "bid1",
						},
					},
					{
						BidPointer: model.BidPointer{
							ExchangeBidID: "bid2",
						},
					},
					{
						BidPointer: model.BidPointer{
							ExchangeBidID: "bid3",
						},
					},
				},
			},
			bidResponse: openrtb.BidResponse{
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								AdMarkup: "https://showcase.com/game/happy-bucket/play?ad_type=banner&adm_id=bid1,bid2,bid3&is_adult=true",
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

			r := &addClickunderADM{}

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.bidResponse, tt.bidResponse)
		})
	}
}
