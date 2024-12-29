package beforeRequestToDSP

import (
	"testing"

	"exchange/internal/services/exchange/model"
	settingsModel "exchange/internal/services/setting/model"
	"pkg/openrtb"
	"pkg/testUtils"
	"pkg/uuid"
)

func Test_clickunder_Apply(t *testing.T) {

	generalSettings := settingsModel.Settings{
		ShowcaseURL: "https://test.com",
	}

	tests := []struct {
		name       string
		req        beforeRequestToDSP
		bidRequest openrtb.BidRequest
		mockValues []string
		wantErr    error
	}{
		{
			name: "1. Пришедший запрос является кликандер запросом, два импрешена",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					IsClickunder: true,
					Settings:     generalSettings,
					PublisherID:  "12345",
				},
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{},
						{},
					},
					App: &openrtb.App{
						Bundle: "appBundle",
					},
				},
			},
			mockValues: []string{"ourSiteID"},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						Interstitial: 1,
						Banner: &openrtb.Banner{
							Formats: []openrtb.Format{
								{
									Width:  730,
									Height: 90,
								},
								{
									Width:  320,
									Height: 50,
								},
								{
									Width:  480,
									Height: 320,
								},
								{
									Width:  300,
									Height: 250,
								},
								{
									Width:  320,
									Height: 90,
								},
								{
									Width:  320,
									Height: 480,
								},
							},
							Width:    0,
							Height:   0,
							Position: openrtb.AdPositionFullscreen,
						},
					},
					{
						Interstitial: 1,
						Banner: &openrtb.Banner{
							Formats: []openrtb.Format{
								{
									Width:  730,
									Height: 90,
								},
								{
									Width:  320,
									Height: 50,
								},
								{
									Width:  480,
									Height: 320,
								},
								{
									Width:  300,
									Height: 250,
								},
								{
									Width:  320,
									Height: 90,
								},
								{
									Width:  320,
									Height: 480,
								},
							},
							Width:    0,
							Height:   0,
							Position: openrtb.AdPositionFullscreen,
						},
					},
				},
				Site: &openrtb.Site{
					Inventory: openrtb.Inventory{
						ID:     "ourSiteID",
						Domain: "12345.test.com",
					},
					Page: "https://test.com",
				},
				App: nil,
			},
			wantErr: nil,
		},
		{
			name: "2. Пришедший запрос является баннерным запросом, два импрешена",
			req: beforeRequestToDSP{
				AuctionDTO: &model.AuctionDTO{
					IsClickunder: false,
				},
				BidRequest: openrtb.BidRequest{
					Impressions: []openrtb.Impression{
						{
							Banner: &openrtb.Banner{
								ID: "bannerID",
							},
						},
						{
							Banner: &openrtb.Banner{
								ID: "bannerID",
							},
						},
					},
					Site: &openrtb.Site{
						Inventory: openrtb.Inventory{
							ID:     "siteID",
							Domain: "vk.com",
						},
					},
					App: &openrtb.App{
						Bundle: "appBundle",
					},
				},
			},
			bidRequest: openrtb.BidRequest{
				Impressions: []openrtb.Impression{
					{
						Banner: &openrtb.Banner{
							ID: "bannerID",
						},
					},
					{
						Banner: &openrtb.Banner{
							ID: "bannerID",
						},
					},
				},
				Site: &openrtb.Site{
					Inventory: openrtb.Inventory{
						ID:     "siteID",
						Domain: "vk.com",
					},
				},
				App: &openrtb.App{
					Bundle: "appBundle",
				},
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			r := &clickunder{}

			uuid.AddMockValues(tt.mockValues...)

			// Применяем метод
			gotErr := r.Apply(&tt.req)
			testUtils.CheckError(t, gotErr, tt.wantErr, false)
			testUtils.CheckStruct(t, tt.req.BidRequest, tt.bidRequest)
		})
	}
}
