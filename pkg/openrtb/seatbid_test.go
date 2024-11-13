package openrtb_test

import (
	"testing"

	"github.com/shopspring/decimal"

	decimalPkg "pkg/decimal"
	"pkg/openrtb"
)

func TestSeatBid_Validate(t *testing.T) {
	testcase := []Testcase{
		{
			Name: "valid",
			Validater: &openrtb.SeatBid{
				Bids: []openrtb.Bid{
					{
						ID:    "1",
						ImpID: "1",
					},
				},
			},
		},
		{
			Name:      "seatbid is missing bids",
			Validater: &openrtb.SeatBid{},
			Err:       openrtb.ErrInvalidSeatBidNoBids,
		},
		{
			Name: "bid is missing ID",
			Validater: &openrtb.SeatBid{
				Bids: []openrtb.Bid{{}},
			},
			Err: openrtb.ErrInvalidBidNoID,
		},
		{
			Name: "bid is missing impression iD",
			Validater: &openrtb.SeatBid{
				Bids: []openrtb.Bid{{ID: "1"}},
			},
			Err: openrtb.ErrInvalidBidNoImpID,
		},
	}

	assertValidate(t, testcase)
}

func TestSeatBid_Unmarshal(t *testing.T) {
	expected := openrtb.SeatBid{
		Bids: []openrtb.Bid{
			{
				ID:        "1",
				ImpID:     "1",
				Price:     decimalPkg.Decimal{Decimal: decimal.NewFromFloat(0.751371)},
				AdID:      "52a5516d29e435137c6f6e74",
				NoticeURL: "http://ads.com/win/112770_1386565997?won=${AUCTION_PRICE}",
				AdMarkup:  "<html/>",
				AdvDomains: []string{
					"ads.com",
				},
				ImageURL:   "http://ads.com/112770_1386565997.jpeg",
				CampaingID: "52a5516d29e435137c6f6e74",
				CreativeID: "52a5516d29e435137c6f6e74_1386565997",
				DealID:     "example_deal",
				Attrs:      []openrtb.CreativeAttribute{},
			},
		},
		Seat:  "1234567-1",
		Group: 1,
	}

	assertEqualJSON(t, "seatbid", &expected)
}
