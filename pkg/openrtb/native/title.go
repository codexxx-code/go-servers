package native

import "encoding/json"

// TitleRequest is the title element of the Native ad.
type TitleRequest struct {
	// Maximum length of the text in the title element.
	//
	// Recommended to be 25, 90, or 140.
	Length int `json:"len"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// TitleResponse is the title element of the Native ad.
type TitleResponse struct {
	// The text associated with the text element.
	Text string `json:"text"`

	// The length of the title being provided.
	//
	// Required: if using assetsurl/dcourl representation.
	//
	// Optional: if using embedded asset representation.
	Length int `json:"len"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
