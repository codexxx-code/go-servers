package native

import "encoding/json"

// Link is used for ‘call to action’ assets, or other links from the Native ad.
type Link struct {
	// Landing URL of the clickable link.
	URL string `json:"url"`

	// List of third-party tracker URLs to be fired on click of
	// the URL.
	ClickTrackers []string `json:"clicktrackers"`

	// Fallback URL for deeplink.
	//
	// To be used if the URL given in url is not supported by
	// the device.
	Fallback string `json:"fallback"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
