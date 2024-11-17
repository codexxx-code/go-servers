package openrtb

import "encoding/json"

// Container for a native impression conforming to the Dynamic Native Ads API.
type Native struct {
	// Request payload complying with the Native Ad Specification.
	//
	// Required.
	Request string `json:"request" bson:"request"`

	// Version of the Dynamic Native Ads API to which request complies.
	//
	// Highly recommended for efficient parsing.
	Version string `json:"ver" bson:"ver"`

	// List of supported API frameworks for this impression.
	//
	// If an API is not explicitly listed, it is assumed not to be supported.
	APIs []APIFramework `json:"api" bson:"api"`

	// Blocked creative attributes.
	BlockedAttrs []CreativeAttribute `json:"battr" bson:"battr"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}
