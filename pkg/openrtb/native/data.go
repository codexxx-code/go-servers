package native

import "encoding/json"

// DataRequest is a non-core element in the Native ad such as Brand Name,
// Ratings, Review Count, Stars, Download count, descriptions etc.
//
// In some cases, additional recommendations are also included in the DataRequest
// Types table.
type DataRequest struct {
	// Type ID of the element supported by the publisher.
	// The publisher can display this information in an appropriate
	// format.
	//
	// Required.
	Type DataType `json:"type"`

	// Maximum length of the text in the element’s response.
	Length int `json:"len"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// DataResponse is a non-core element in the Native ad such as Brand Name,
// Ratings, Review Count, Stars, Download count, descriptions etc.
//
// In some cases, additional recommendations are also included in the DataResponse
// Types table.
type DataResponse struct {
	// Type ID of the element supported by the publisher.
	// The publisher can display this information in an appropriate
	// format.
	//
	// Required.
	Type DataType `json:"type"`

	// Maximum length of the text in the element’s response.
	//
	// Required: for assetsurl/dcourl responses.
	//
	// Not required: for embedded asset responses.
	Length int `json:"len"`

	// The optional formatted string name of the data type to be
	// displayed.
	//
	// Deprecated: since version 1.2.
	Label string `json:"label"`

	// The formatted string of data to be displayed.
	//
	// Can contain a formatted value such as "5 stars" or "$10"
	// or "3.4 stars out of 5".
	Value string `json:"value"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
