package openrtb

import "encoding/json"

// Details for a banner impression (incl. in-banner video) or video companion ad.
type Banner struct {
	// Array of format objects representing the banner sizes permitted.
	//
	// If none are specified, then use of the h and w attributes is highly recommended.
	Formats []Format `json:"format" bson:"format"`

	// Exact width in device independent pixels (DIPS).
	//
	// Recommended if no format objects are specified.
	Width int `json:"w" bson:"w"`

	// Exact height in device independent pixels (DIPS).
	//
	// Recommended if no format objects are specified
	Height int `json:"h" bson:"h"`

	// Maximum width in device independent pixels (DIPS).
	// Deprecated: deprecated in favor of the format array
	WidthMax int `json:"wmax" bson:"wmax"`

	// Maximum height in device independent pixels (DIPS).
	//
	// Deprecated: deprecated in favor of the format array.
	HeightMax int `json:"hmax" bson:"hmax"`

	// Minimum width in device independent pixels (DIPS).
	//
	// Deprecated: deprecated in favor of the format array.
	WidthMin int `json:"wmin" bson:"wmin"`

	// Minimum height in device independent pixels (DIPS).
	//
	// Deprecated: deprecated in favor of the format array.
	HeightMin int `json:"hmin" bson:"hmin"`

	// Blocked banner ad types.
	BlockedTypes []BannerType `json:"btype" bson:"btype"`

	// Blocked creative attributes.
	BlockedAttrs []CreativeAttribute `json:"battr" bson:"battr"`

	// Ad position on screen.
	Position AdPosition `json:"pos" bson:"pos"`

	// Content MIME types supported. Popular MIME types may include
	// “application/x-shockwave-flash”, “image/jpg”, and “image/gif”.
	MIMEs []string `json:"mimes" bson:"mimes"`

	// Indicates if the banner is in the top frame as opposed to an iframe, where:
	//    0 = no;
	//    1 = yes.
	TopFrame int `json:"topframe" bson:"topframe"`

	// Directions in which the banner may expand.
	ExpDirs []ExpDir `json:"expdir" bson:"expdir"`

	// List of supported API frameworks for this impression.
	//
	// If an API is not explicitly listed, it is assumed not to be supported.
	APIs []APIFramework `json:"api" bson:"api"`

	// Unique identifier for this banner object.
	//
	// Recommended when Banner objects are used with a Video object to represent an array
	// of companion ads. Values usually start at 1 and increase with each object;
	// should be unique within an impression.
	ID string `json:"id" bson:"id"`

	// Relevant only for Banner objects used with a Video object in an array of companion ads.
	// Indicates the companion banner rendering mode relative to the associated video, where:
	//   0 = concurrent;
	//   1 = end-card.
	VCM int `json:"vcm" bson:"vcm"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}
