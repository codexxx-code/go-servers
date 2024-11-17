package openrtb

import "encoding/json"

// ThirdParty abstract attributes.
type ThirdParty struct {
	// Content producer or originator ID. Useful if content is syndicated and may be
	// posted on a site using embed tags.
	ID string `json:"id" bson:"id"`

	// Content producer or originator name (e.g., “Warner Bros”)
	Name string `json:"name" bson:"name"`

	// Array of IAB content categories that describe the content producer.
	Categories []ContentCategory `json:"cat" bson:"cat"`

	// Highest level domain of the content producer (e.g., “producer.com”).
	Domain string `json:"domain" bson:"domain"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}

// Entity that controls the content of and distributes the site or app.
type Publisher ThirdParty

// Producer of the content; not necessarily the publisher (e.g., syndication).
type Producer ThirdParty
