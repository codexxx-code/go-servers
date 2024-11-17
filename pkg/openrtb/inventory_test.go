package openrtb_test

import (
	"testing"

	"pkg/openrtb"
)

func TestApp_Unmarshal(t *testing.T) {
	expected := openrtb.App{
		Inventory: openrtb.Inventory{
			ID:   "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfIUDA",
			Name: "Yahoo Weather",
			Categories: []openrtb.ContentCategory{
				openrtb.ContentCategoryScience,
				openrtb.ContentCategoryWeather,
			},
			Publisher: &openrtb.Publisher{
				ID:     "agltb3B1Yi1pbmNyDAsSA0FwcBiJkfTUCV",
				Name:   "yahoo",
				Domain: "www.yahoo.com",
			},
		},
		Version:  "1.0.2",
		Bundle:   "628677149",
		StoreURL: "https://itunes.apple.com/id628677149",
	}

	assertEqualJSON(t, "app", &expected)
}

func TestSite_Unmarshal(t *testing.T) {
	expected := openrtb.Site{
		Inventory: openrtb.Inventory{
			ID:     "102855",
			Domain: "http://www.usabarfinder.com",
			Categories: []openrtb.ContentCategory{
				openrtb.ContentCategoryAdvertising,
			},
			Publisher: &openrtb.Publisher{
				ID:     "8953",
				Name:   "local.com",
				Domain: "local.com",
			},
		},
		Page: "http://eas.usabarfinder.com/eas?cu=13824;cre=mu;target=_blank",
	}

	assertEqualJSON(t, "site", &expected)
}
