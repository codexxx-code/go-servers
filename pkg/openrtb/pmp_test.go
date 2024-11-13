package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestPMP_Validate(t *testing.T) {
	testcases := []Testcase{
		{
			Name: "valid",
			Validater: &openrtb.PMP{
				Deals: []openrtb.Deal{
					{
						ID: "id",
					},
				},
			},
		},
		{
			Name: "no valid",
			Validater: &openrtb.PMP{
				Deals: []openrtb.Deal{{}},
			},
			Err: openrtb.ErrInvalidDealNoID,
		},
	}

	assertValidate(t, testcases)
}

func TestPMP_Unmarshal(t *testing.T) {
	expected := openrtb.PMP{
		Private: 1,
		Deals: []openrtb.Deal{
			{
				ID:               "DX-1985-010A",
				BidFloor:         2.5,
				BidFloorCurrency: "RUB",
				AuctionType:      2,
			},
			{
				ID:               "DX-1986-010A",
				BidFloor:         2.6,
				BidFloorCurrency: "USD",
			},
		},
	}

	assertEqualJSON(t, "pmp", &expected)
}
