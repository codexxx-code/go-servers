package native

import "encoding/json"

// ImageRequest is an image element in native ad, such as icons, main image, etc.
type ImageRequest struct {
	// Type ID of the image element supported by the publisher.
	// The publisher can display this information in an appropriate
	// format.
	Type ImageType `json:"type"`

	// Width of the image in pixels.
	Width int `json:"w"`

	// The minimum requested width of the image in pixels.
	//
	// This option should be used for any rescaling of images by
	// the client.
	// Either Width or WidthMin should be transmitted.
	// If only Width is included, it should be considered an exact
	// requirement.
	//
	// Recommended.
	WidthMin int `json:"wmin"`

	// Height of the image in pixels.
	Height int `json:"h"`

	// The minimum requested height of the image in pixels.
	//
	// This option should be used for any rescaling of images by
	// the client.
	// Either Height or HeightMin should be transmitted.
	// If only Height is included, it should be considered an exact
	// requirement.
	//
	// Recommended.
	HeightMin int `json:"hmin"`

	// Whitelist of content MIME types supported.
	//
	// If blank, assume all types are allowed.
	MIMEs []string `json:"mimes"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}

// ImageResponse is an image element in native ad, such as icons, main image, etc.
type ImageResponse struct {
	// Type ID of the image element supported by the publisher.
	// The publisher can display this information in an appropriate
	// format.
	//
	// Required: for assetsurl or dcourl responses.
	//
	// Not required: for embedded asset responses.
	Type ImageType `json:"type"`

	// URL of the image asset
	URL string `json:"url"`

	// Width of the image in pixels.
	//
	// Recommended: for embedded asset responses.
	//
	// Required: for assetsurl/dcourlresponses if multiple assets
	// of same type submitted.
	Width int `json:"w"`

	// Height of the image in pixels.
	//
	// Recommended: for embedded asset responses.
	//
	// Required: for assetsurl/dcourlresponses if multiple assets
	// of same type submitted.
	Height int `json:"h"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty"`
}
