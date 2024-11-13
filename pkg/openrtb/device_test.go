package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestDevice_Unmarshal(t *testing.T) {
	expected := openrtb.Device{
		DNT:       0,
		UserAgent: "Mozilla/5.0 (iPhone; CPU iPhone OS 6_1 like Mac OS X) AppleWebKit/534.46 (KHTML, like Gecko) Version/5.1 Mobile/9A334 Safari/7534.48.3",
		IP:        "123.145.167.189",
		Geo: &openrtb.Geo{
			Country:    "USA",
			Latitude:   35.012345,
			Longtitude: -115.12345,
			City:       "Los Angeles",
			Metro:      "803",
			Region:     "CA",
			ZIP:        "90049",
		},
		PIDSHA1:        "AA000DFE74168477C70D291f574D344790E0BB11",
		PIDMD5:         "AA003EABFB29E6F759F3BDAB34E50BB11",
		Carrier:        "310-410",
		Language:       "en",
		Make:           "Apple",
		Model:          "iPhone",
		OS:             "iOS",
		OSVersion:      "6.1",
		JS:             1,
		ConnectionType: openrtb.ConnectionTypeCell,
		MCCMNC:         "722-341",
		DeviceType:     openrtb.DeviceTypeMobile,
	}

	assertEqualJSON(t, "device", &expected)
}
