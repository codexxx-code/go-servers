package openrtb_test

import (
	"testing"

	"github.com/shopspring/decimal"

	decimalPkg "pkg/decimal"
	"pkg/openrtb"
)

func TestBid_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name:      "valid",
			Validater: &openrtb.Bid{ID: "id", ImpID: "id"},
		},
		{
			Name:      "bid is missing ID",
			Validater: &openrtb.Bid{},
			Err:       openrtb.ErrInvalidBidNoID,
		},
		{
			Name:      "bid is missing impression ID",
			Validater: &openrtb.Bid{ID: "id"},
			Err:       openrtb.ErrInvalidBidNoImpID,
		},
	}

	assertValidate(t, testcases)
}

func TestBid_Unmarshal(t *testing.T) {
	expected := openrtb.Bid{
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
	}

	assertEqualJSON(t, "bid", &expected)
}
