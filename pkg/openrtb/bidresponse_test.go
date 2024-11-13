package openrtb_test

import (
	"testing"

	"github.com/shopspring/decimal"

	decimalPkg "pkg/decimal"
	"pkg/openrtb"
	"pkg/testUtils"
)

func TestBidReponse_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name: "valid",
			Validater: &openrtb.BidResponse{
				ID: "id",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								ID:    "id",
								ImpID: "impid",
							},
						},
					},
				},
			},
		},
		{
			Name:      "bidresponse is missing ID",
			Validater: &openrtb.BidResponse{},
			Err:       openrtb.ErrInvalidBidResponseNoID,
		},
		{
			Name: "seatbid is missing bids",
			Validater: &openrtb.BidResponse{
				ID: "id",
			},
			Err: openrtb.ErrInvalidSeatBidNoBids,
		},
		{
			Name: "bid is missing ID",
			Validater: &openrtb.BidResponse{
				ID: "id",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{{}},
					},
				},
			},
			Err: openrtb.ErrInvalidBidNoID,
		},
		{
			Name: "bid is missing impression ID",
			Validater: &openrtb.BidResponse{
				ID: "id",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{{ID: "id"}},
					},
				},
			},
			Err: openrtb.ErrInvalidBidNoImpID,
		},
	}

	assertValidate(t, testcases)
}

func TestBidResponse_Unmarshal(t *testing.T) {
	testcases := []struct {
		file string
		exp  openrtb.BidResponse
	}{
		{
			file: "bidresponse_single",
			exp: openrtb.BidResponse{
				ID:       "BID-4-ZIMP-4b309eae-504a-4252-a8a8-4c8ceee9791a",
				Currency: "USD",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								ID:         "32a69c6ba388f110487f9d1e63f77b22d86e916b",
								ImpID:      "32a69c6ba388f110487f9d1e63f77b22d86e916b",
								Price:      decimalPkg.Decimal{Decimal: decimal.NewFromFloat(0.065445)},
								AdID:       "529833ce55314b19e8796116",
								NoticeURL:  "http://ads.com/win/529833ce55314b19e8796116?won=${auction_price}",
								AdMarkup:   "<iframe src=\"foo.bar\"/>",
								CampaingID: "529833ce55314b19e8796116",
								CreativeID: "529833ce55314b19e8796116_1385706446",
							},
						},
						Seat: "772",
					},
				},
			},
		},
		{
			file: "bidresponse_multi",
			exp: openrtb.BidResponse{
				ID:       "BID-4-ZIMP-4b309eae-504a-4252-a8a8-4c8ceee9791a",
				Currency: "USD",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								ID:         "24195efda36066ee21f967bc1de14c82db841f07",
								ImpID:      "24195efda36066ee21f967bc1de14c82db841f07",
								Price:      decimalPkg.Decimal{Decimal: decimal.NewFromFloat(1.028428)},
								AdID:       "52a12b5955314b7194a4c9ff",
								NoticeURL:  "http://ads.com/win/52a12b5955314b7194a4c9ff?won=${AUCTION_PRICE}",
								AdMarkup:   "<iframe />",
								AdvDomains: []string{"ads.com"},
								CampaingID: "52a12b5955314b7194a4c9ff",
								CreativeID: "52a12b5955314b7194a4c9ff_1386294105",
								DealID:     "DX-1985-010A",
							},
						},
						Seat: "42",
					},
					{
						Bids: []openrtb.Bid{
							{
								ID:         "24195efda36066ee21f967bc1de14c82db841f08",
								ImpID:      "24195efda36066ee21f967bc1de14c82db841f08",
								Price:      decimalPkg.Decimal{Decimal: decimal.NewFromFloat(0.04958)},
								AdID:       "527c9fdd55314ba06815f25e",
								NoticeURL:  "http://ads.com/win/527c9fdd55314ba06815f25e?won=${AUCTION_PRICE}",
								AdMarkup:   "<iframe />",
								AdvDomains: []string{"ads.com"},
								CampaingID: "527c9fdd55314ba06815f25e",
								CreativeID: "527c9fdd55314ba06815f25e_1383899102",
							},
						},
						Seat: "772",
					},
				},
			},
		},
		{
			file: "bidresponse_pmp",
			exp: openrtb.BidResponse{
				ID:       "1234567890",
				BidID:    "abc1123",
				Currency: "USD",
				SeatBids: []openrtb.SeatBid{
					{
						Bids: []openrtb.Bid{
							{
								ID:        "1",
								ImpID:     "102",
								Price:     decimalPkg.Decimal{Decimal: testUtils.IgnoreError(decimal.NewFromString("5.00"))},
								DealID:    "ABC-1234-6789",
								NoticeURL: "http: //adserver.com/winnotice?impid=102",
								AdvDomains: []string{
									"advertiserdomain.com",
								},
								ImageURL:   "http: //adserver.com/pathtosampleimage",
								CampaingID: "campaign111",
								CreativeID: "creative112",
								AdID:       "314",
								Attrs: []openrtb.CreativeAttribute{
									openrtb.CreativeAttributeAudioAdAutoPlay,
									openrtb.CreativeAttributeAudioAdUserInitiated,
									openrtb.CreativeAttributeExpandableAuto,
									openrtb.CreativeAttributeExpandableUserInitiatedClick,
								},
							},
						},
						Seat: "512",
					},
				},
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
