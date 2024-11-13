package openrtb_test

import (
	"testing"

	"github.com/shopspring/decimal"

	decimalPkg "pkg/decimal"
	"pkg/openrtb"
)

func TestImpression_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name:      "valid",
			Validater: &openrtb.Impression{ID: "id"},
		},
		{
			Name:      "missing id",
			Validater: &openrtb.Impression{},
			Err:       openrtb.ErrInvalidImpNoID,
		},
		{
			Name: "multiple assets",
			Validater: &openrtb.Impression{
				ID:     "id",
				Banner: &openrtb.Banner{},
				Video:  &openrtb.Video{},
				Audio:  &openrtb.Audio{},
				Native: &openrtb.Native{},
			},
			Err: openrtb.ErrInvalidMultiAssets,
		},
		{
			Name: "video not valid",
			Validater: &openrtb.Impression{
				ID:    "id",
				Video: &openrtb.Video{},
			},
			Err: openrtb.ErrInvalidVideoNoMIMEs,
		},
		{
			Name: "audio not valid",
			Validater: &openrtb.Impression{
				ID:    "id",
				Audio: &openrtb.Audio{},
			},
			Err: openrtb.ErrInvalidAudioNoMIMEs,
		},
	}

	assertValidate(t, testcases)
}

func TestImpression_Unmarshal(t *testing.T) {
	expected := openrtb.Impression{
		ID: "1",
		Banner: &openrtb.Banner{
			Height:   250,
			Width:    300,
			Position: openrtb.AdPositionUnknown,
		},
		BidFloor:         decimalPkg.Decimal{Decimal: decimal.NewFromFloat(0.03)},
		BidFloorCurrency: "USD",
		PMP: &openrtb.PMP{
			Private: 1,
			Deals: []openrtb.Deal{
				{
					ID:               "DX-1985-010A",
					BidFloor:         2.5,
					BidFloorCurrency: "USD",
					AuctionType:      2,
				},
			},
		},
	}

	assertEqualJSON(t, "impression", &expected)
}
