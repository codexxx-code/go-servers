package native

import "encoding/json"

// EventTrackerRequest is a specifies the types of events the bidder can
// request to be tracked in the bid response, and which types of tracking are
// available for each event type, and is included as an array in the request.
type EventTrackerRequest struct {
	// Type of event available for tracking.
	Type EventType `json:"event"`

	// Array of the types of tracking available for the given event.
	Methods []EventTrackingMethod `json:"methods"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// EventTrackerResponse is a specifies the types of events the bidder wishes to
// track and the URLs/information to track them.
//
// Bidder must only respond with methods indicated as available in the request.
//
// Note that most javascript trackers expect to be loaded at impression time,
// so it’s not generally recommended for the buyer to respond with javascript
// trackers on other events, but the appropriateness of this is up to each buyer.
type EventTrackerResponse struct {
	// Type of event to track.
	Type EventType `json:"event"`

	// Type of tracking requested
	Method EventTrackingMethod `json:"method"`

	// The URL of the image or js.
	//
	// Required: for image or js, optional for custom.
	URL string `json:"url"`

	// To be agreed individually with the exchange, an array of
	// key:value objects for custom tracking, for example
	// the account number of the DSP with a tracking company.
	//
	// Type is a key-value object (not specified if value is
	// string-only).
	CustomData json.RawMessage `json:"customdata"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
