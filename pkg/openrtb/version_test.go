package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestHeaderStringer(t *testing.T) {
	h := openrtb.OpenRTBVersionHeader
	assertEqualValues(t, "x-openrtb-version", h.String())
}

func TestVersionStringer(t *testing.T) {
	v := openrtb.OpenRTBVersion25
	assertEqualValues(t, "2.5", v.String())
}
