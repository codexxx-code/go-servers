package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestBanner_Unmarshal(t *testing.T) {
	expected := openrtb.Banner{
		Width:    728,
		Height:   90,
		Position: openrtb.AdPositionAboveFold,
		BlockedTypes: []openrtb.BannerType{
			openrtb.BannerTypeFrame,
		},
		BlockedAttrs: []openrtb.CreativeAttribute{
			openrtb.CreativeAttributeWindowsDialogOrAlert,
		},
		APIs: []openrtb.APIFramework{
			openrtb.APIFrameworkMRAID1,
		},
		VCM: 1,
	}

	assertEqualJSON(t, "banner", &expected)
}
