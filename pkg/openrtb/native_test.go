package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestNative_Unmarshal(t *testing.T) {
	expected := openrtb.Native{
		Request: "PAYLOAD",
		Version: "2",
		APIs: []openrtb.APIFramework{
			openrtb.APIFrameworkVPAID1,
			openrtb.APIFrameworkVPAID2,
		},
		BlockedAttrs: []openrtb.CreativeAttribute{
			openrtb.CreativeAttributeUserInitiated,
			openrtb.CreativeAttributeWindowsDialogOrAlert,
		},
	}

	assertEqualJSON(t, "native", &expected)
}
