package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestContent_Unmarshal(t *testing.T) {
	expected := openrtb.Content{
		ID:    "1",
		Title: "Video",
		Producer: &openrtb.Producer{
			ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
			Name:   "yahoo",
			Domain: "www.yahoo.com",
		},
	}

	assertEqualJSON(t, "content", &expected)
}
