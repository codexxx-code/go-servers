package openrtb_test

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"

	decimalPkg "pkg/decimal"
	"pkg/openrtb"
)

func TestBidRequest_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name: "valid",
			Validater: &openrtb.BidRequest{
				ID: "id",
				Impressions: []openrtb.Impression{{
					ID: "id",
				}},
			},
		},
		{
			Name:      "missing ID",
			Validater: &openrtb.BidRequest{},
			Err:       openrtb.ErrInvalidRequestNoID,
		},
		{
			Name: "missing impressions",
			Validater: &openrtb.BidRequest{
				ID: "id",
			},
			Err: openrtb.ErrInvalidRequestNoImps,
		},
		{
			Name: "multi inventory",
			Validater: &openrtb.BidRequest{
				ID: "id",
				Impressions: []openrtb.Impression{{
					ID: "id",
				}},
				Site: &openrtb.Site{},
				App:  &openrtb.App{},
			},
			Err: openrtb.ErrInvalidRequestMultiInv,
		},
		{
			Name: "not valid impression",
			Validater: &openrtb.BidRequest{
				ID: "id",
				Impressions: []openrtb.Impression{{
					ID: "",
				}},
			},
			Err: openrtb.ErrInvalidImpNoID,
		},
	}

	assertValidate(t, testcases)
}

func TestBidRequest_Unmarshal(t *testing.T) {
	testcases := []struct {
		file string
		exp  openrtb.BidRequest
	}{
		{
			file: "bidrequest_banner",
			exp: openrtb.BidRequest{
				ID:          "1234534625254",
				AuctionType: 2,
				TimeMax:     120,
				Impressions: []openrtb.Impression{
					{
						ID:               "1",
						Secure:           1,
						BidFloorCurrency: "USD",
						Banner: &openrtb.Banner{
							Width:    300,
							Height:   250,
							Position: openrtb.AdPositionAboveFold,
							BlockedAttrs: []openrtb.CreativeAttribute{
								openrtb.CreativeAttributeUserInitiated,
							},
						},
					},
				},
				BlockedAdvDomains: []string{
					"company1.com",
					"company2.com",
				},
				Site: &openrtb.Site{
					Inventory: openrtb.Inventory{
						ID:     "234563",
						Name:   "Site ABCD",
						Domain: "siteabcd.com",
						Categories: []openrtb.ContentCategory{
							openrtb.ContentCategoryAutoParts,
							openrtb.ContentCategoryAutoRepair,
						},
						PrivacyPolicy: 1,
						Publisher: &openrtb.Publisher{
							ID:   "pub12345",
							Name: "Publisher A",
						},
						Content: &openrtb.Content{
							Keywords: "keyword a,keyword b,keyword c",
						},
					},
					Page:     "http://siteabcd.com/page.htm",
					Refferer: "http://referringsite.com/referringpage.htm",
				},
				Device: &openrtb.Device{
					IP:           "64.124.253.1",
					UserAgent:    "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20110319 Firefox/3.6.16",
					OS:           "OS X",
					FlashVersion: "10.1",
					JS:           1,
				},
				User: &openrtb.User{
					ID:       "45asdf987656789adfad4678rew656789",
					BuyerUID: "5df678asd8987656asdf78987654",
				},
				Test: 1,
			},
		},
		{
			file: "bidrequest_video",
			exp: openrtb.BidRequest{
				ID:          "0123456789ABCDEF0123456789ABCDEF",
				AuctionType: 2,
				TimeMax:     120,
				Impressions: []openrtb.Impression{
					{
						ID:               "1",
						BidFloorCurrency: "USD",
						PMP: &openrtb.PMP{
							Private: 1,
							Deals: []openrtb.Deal{
								{
									ID:               "1452f.eadb4.7aaa",
									BidFloor:         5.3,
									BidFloorCurrency: "USD",
									AuctionType:      1,
									Ext:              json.RawMessage([]byte(`{"priority":1,"wadvs":[]}`)),
								},
							},
						},
						Video: &openrtb.Video{
							MIMEs: []string{
								"video/x-flv",
								"video/mp4",
								"application/x-shockwave-flash",
								"application/javascript",
							},
							APIs: []openrtb.APIFramework{
								openrtb.APIFrameworkVPAID1,
								openrtb.APIFrameworkVPAID2,
							},
							BlockedAttrs: []openrtb.CreativeAttribute{
								openrtb.CreativeAttributeUserInitiated,
								openrtb.CreativeAttributeWindowsDialogOrAlert,
							},
							BoxingAllowed: 1,
							Delivery: []openrtb.ContentDelivery{
								openrtb.ContentDeliveryProgressive,
							},
							Width:       640,
							Height:      480,
							MaxBitrate:  1500,
							MinBitrate:  300,
							MaxDuration: 30,
							MinDuration: 5,
							Linearity:   openrtb.VideoLinearityLinear,
							PlaybackMethods: []openrtb.VideoPlayback{
								openrtb.VideoPlaybackPageLoadSoundOn,
							},
							Position: openrtb.AdPositionAboveFold,
							Protocols: []openrtb.Protocol{
								openrtb.ProtocolVAST2,
								openrtb.ProtocolVAST3,
							},
							Sequence:   1,
							StartDelay: openrtb.StartDelayPreRoll,
						},
					},
				},
				Site: &openrtb.Site{
					Inventory: openrtb.Inventory{
						ID:     "1345135123",
						Name:   "Site ABCD",
						Domain: "siteabcd.com",
						Categories: []openrtb.ContentCategory{
							openrtb.ContentCategoryAutoParts,
							openrtb.ContentCategoryAutoRepair,
						},
						PrivacyPolicy: 1,
						Publisher: &openrtb.Publisher{
							ID:   "pub12345",
							Name: "Publisher A",
						},
						Content: &openrtb.Content{
							ID: "1234567",
							Categories: []openrtb.ContentCategory{
								openrtb.ContentCategoryAutoRepair,
							},
							Episode:  23,
							Season:   "2",
							Title:    "Car Show",
							Series:   "All About Cars",
							Keywords: "keyword a,keyword b,keyword c",
						},
					},
					Page:     "http://siteabcd.com/page.htm",
					Refferer: "http://referringsite.com/referringpage.htm",
				},
				Device: &openrtb.Device{
					IP:           "64.124.253.1",
					UserAgent:    "Mozilla/5.0 (Mac; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20140420 Firefox/3.6.16",
					OS:           "OS X",
					FlashVersion: "10.1",
					JS:           1,
				},
				User: &openrtb.User{
					ID:       "456789876567897654678987656789",
					BuyerUID: "545678765467876567898765678987654",
					Data: []openrtb.Data{
						{
							ID:   "6",
							Name: "Data Provider 1",
							Segment: []openrtb.Segment{
								{
									ID:   "12341318394918",
									Name: "auto intenders",
								},
								{
									ID:   "1234131839491234",
									Name: "auto enthusiasts",
								},
							},
						},
					},
				},
			},
		},
		{
			file: "bidrequest_audio",
			exp: openrtb.BidRequest{
				ID:          "0123456789ABCDEF0123456789ABCDEF",
				AuctionType: 2,
				TimeMax:     120,
				Impressions: []openrtb.Impression{
					{
						ID:               "1",
						BidFloorCurrency: "USD",
						PMP: &openrtb.PMP{
							Private: 1,
							Deals: []openrtb.Deal{
								{
									ID:               "1452f.eadb4.7aaa",
									BidFloor:         5.3,
									BidFloorCurrency: "USD",
									AuctionType:      1,
									Ext:              json.RawMessage([]byte(`{"priority":1,"wadvs":[]}`)),
								},
							},
						},
						Audio: &openrtb.Audio{
							MIMEs:       []string{"audio/mp4"},
							MinDuration: 5,
							MaxDuration: 30,
							Protocols: []openrtb.Protocol{
								openrtb.ProtocolDAAST1,
								openrtb.ProtocolDAAST1Wrapper,
							},
							BlockedAttrs: []openrtb.CreativeAttribute{
								openrtb.CreativeAttributeUserInitiated,
								openrtb.CreativeAttributeWindowsDialogOrAlert,
							},
							Sequence:    0,
							MaxExtended: 30,
							MinBitrate:  300,
							MaxBitrate:  1500,
							Delivery: []openrtb.ContentDelivery{
								openrtb.ContentDeliveryProgressive,
							},
							CompanionAds: []openrtb.Banner{
								{
									ID:       "1234567893-1",
									Width:    300,
									Height:   250,
									Position: openrtb.AdPositionAboveFold,
									BlockedAttrs: []openrtb.CreativeAttribute{
										openrtb.CreativeAttributeUserInitiated,
										openrtb.CreativeAttributeWindowsDialogOrAlert,
									},
									ExpDirs: []openrtb.ExpDir{
										openrtb.ExpDirRight,
										openrtb.ExpDirDown,
									},
								},
							},
							APIs: []openrtb.APIFramework{
								openrtb.APIFrameworkVPAID1,
								openrtb.APIFrameworkVPAID2,
							},
							CompanionTypes: []openrtb.CompanionType{
								openrtb.CompanionTypeStatic,
								openrtb.CompanionTypeHTML,
							},
						},
					},
				},
				Site: &openrtb.Site{
					Inventory: openrtb.Inventory{
						ID:     "1345135123",
						Name:   "Site ABCD",
						Domain: "siteabcd.com",
						Categories: []openrtb.ContentCategory{
							openrtb.ContentCategoryAutoParts,
							openrtb.ContentCategoryAutoRepair,
						},
						PrivacyPolicy: 1,
						Publisher: &openrtb.Publisher{
							ID:   "pub12345",
							Name: "Publisher A",
						},
						Content: &openrtb.Content{
							ID: "1234567",
							Categories: []openrtb.ContentCategory{
								openrtb.ContentCategoryAutoRepair,
							},
							Episode:  23,
							Season:   "2",
							Title:    "Car Show",
							Series:   "All About Cars",
							Keywords: "keyword a,keyword b,keyword c",
						},
					},
					Page:     "http://siteabcd.com/page.htm",
					Refferer: "http://referringsite.com/referringpage.htm",
				},
				Device: &openrtb.Device{
					IP:           "64.124.253.1",
					UserAgent:    "Mozilla/5.0 (Mac; U; Intel Mac OS X 10.6; en-US; rv:1.9.2.16) Gecko/20140420 Firefox/3.6.16",
					OS:           "OS X",
					FlashVersion: "10.1",
					JS:           1,
				},
				User: &openrtb.User{
					ID:       "456789876567897654678987656789",
					BuyerUID: "545678765467876567898765678987654",
					Data: []openrtb.Data{
						{
							ID:   "6",
							Name: "Data Provider 1",
							Segment: []openrtb.Segment{
								{
									ID:   "12341318394918",
									Name: "auto intenders",
								},
								{
									ID:   "1234131839491234",
									Name: "auto enthusiasts",
								},
							},
						},
					},
				},
			},
		},
		{
			file: "bidrequest_native",
			exp: openrtb.BidRequest{
				ID:          "80ce30c53c16e6ede735f123ef6e32361bfc7b22",
				AuctionType: 1,
				Currencies:  []string{"USD"},
				Impressions: []openrtb.Impression{
					{
						ID:               "1",
						BidFloor:         decimalPkg.Decimal{Decimal: decimal.NewFromFloat(0.03)},
						BidFloorCurrency: "USD",
						Native: &openrtb.Native{
							Request: "...Native Spec request as an encoded string...",
							Version: "1.0",
							APIs: []openrtb.APIFramework{
								openrtb.APIFrameworkMRAID1,
							},
							BlockedAttrs: []openrtb.CreativeAttribute{
								openrtb.CreativeAttributeUserInitiated,
								openrtb.CreativeAttributeWindowsDialogOrAlert,
							},
						},
					},
				},
				Site: &openrtb.Site{
					Inventory: openrtb.Inventory{
						ID: "102855",
						Categories: []openrtb.ContentCategory{
							openrtb.ContentCategoryAdvertising,
						},
						Domain: "www.foobar.com",
						Publisher: &openrtb.Publisher{
							ID:   "8953",
							Name: "foobar.com",
							Categories: []openrtb.ContentCategory{
								openrtb.ContentCategoryAdvertising,
							},
							Domain: "foobar.com",
						},
					},
					Page: "http://www.foobar.com/1234.html ",
				},
				Device: &openrtb.Device{
					UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_6_8) AppleWebKit/537.13 (KHTML, like Gecko) Version/5.1.7 Safari/534.57.2",
					IP:        "123.145.167.10",
				},
				User: &openrtb.User{
					ID: "55816b39711f9b5acf3b90e313ed29e51665623f",
				},
				Seats:        []string{"771", "772"},
				BlockedSeats: []string{"800", "773"},
			},
		},
	}

	for _, tc := range testcases {
		tc := tc
		t.Run(tc.file, func(t *testing.T) {
			t.Parallel()
			assertEqualJSON(t, tc.file, &tc.exp)
		})
	}
}
