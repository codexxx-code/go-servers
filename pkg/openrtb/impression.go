package openrtb

import (
	"encoding/json"
	"errors"

	"pkg/decimal"
)

// Validation errors.
var (
	ErrInvalidImpNoID     = errors.New("impression has no ID")
	ErrInvalidMultiAssets = errors.New("impression has multiple assets")
)

// Container for the description of a specific impression.
type Impression struct {
	// A unique identifier for this impression within the context of the bid request
	// (typically, starts with 1 and increments).
	//
	// Required.
	ID string `json:"id" bson:"id"`

	// An array of Metric object.
	Metric []Metric `json:"metric" bson:"metric"`

	// A Banner object.
	//
	// Required if this impression is offered as a banner ad opportunity.
	Banner *Banner `json:"banner" bson:"banner"`

	// A Video object.
	//
	// Required if this impression is offered as a video ad opportunity.
	Video *Video `json:"video" bson:"video"`

	// An Audio object.
	//
	// Required if this impression is offered as an audio ad opportunity.
	Audio *Audio `json:"audio" bson:"audio"`

	// A Native object.
	//
	// Required if this impression is offered as a native ad opportunity.
	Native *Native `json:"native" bson:"native"`

	// A PMP object containing any private marketplace deals in effect for this impression.
	PMP *PMP `json:"pmp" bson:"pmp"`

	// Name of ad mediation partner, SDK technology, or player responsible for rendering
	// ad (typically video or mobile). Used by some ad servers to customize ad code by partner.
	//
	// Recommended for video and/or apps.
	DisplayManager string `json:"displaymanager" bson:"displaymanager"`

	// Version of ad mediation partner, SDK technology, or player responsible for rendering
	// ad (typically video or mobile). Used by some ad servers to customize ad code by partner.
	//
	// Recommended for video and/or apps.
	DisplayManagerVersion string `json:"displaymanagerver" bson:"displaymanagerver"`

	//    1 = the ad is interstitial or full screen;
	//    0 = not interstitial.
	//
	// Default 0.
	Interstitial int `json:"instl" bson:"instl"`

	// Identifier for specific ad placement or ad tag that was used to initiate the auction.
	// This can be useful for debugging of any issues, or for optimization by the buyer.
	TagID string `json:"tagid" bson:"tagid"`

	// Minimum bid for this impression expressed in CPM.
	//
	// Default 0.0.
	BidFloor decimal.Decimal `json:"bidfloor" bson:"bidfloor"`

	// Currency specified using ISO-4217 alpha codes. This may be different from bid currency
	// returned by bidder if this is allowed by the exchange.
	//
	// Default USD.
	BidFloorCurrency string `json:"bidfloorcur" bson:"bidfloorcur"`

	// Indicates the type of browser opened upon clicking the creative in an app, where:
	//    0 = embedded;
	//    1 = native.
	// Note that the Safari View Controller in iOS 9.x devices is considered a native browser
	// for purposes of this attribute.
	ClickBrowser int `json:"clickbrowser" bson:"clickbrowser"`

	// Flag to indicate if the impression requires secure HTTPS URL creative assets and markup,
	// where:
	//    0 = non-secure;
	//    1 = secure.
	// If omitted, the secure state is unknown, but non-secure HTTP support can be assumed.
	Secure int `json:"secure" bson:"secure"`

	// Array of exchange-specific names of supported iframe busters.
	IFrameBusters []string `json:"iframebuster" bson:"iframebuster"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}

func (imp *Impression) assetCount() int {
	var n int
	if imp.Banner != nil {
		n++
	}
	if imp.Video != nil {
		n++
	}
	if imp.Audio != nil {
		n++
	}
	if imp.Native != nil {
		n++
	}
	return n
}

// Validate the impression object.
func (imp *Impression) Validate() error {
	if imp.ID == "" {
		return ErrInvalidImpNoID
	} else if count := imp.assetCount(); count > 1 {
		return ErrInvalidMultiAssets
	}

	if imp.Video != nil {
		if err := imp.Video.Validate(); err != nil {
			return err
		}
	}

	if imp.Audio != nil {
		if err := imp.Audio.Validate(); err != nil {
			return err
		}
	}

	return nil
}

type jsonImpression Impression

// UnmarshalJSON custom unmarshalling.
func (imp *Impression) UnmarshalJSON(data []byte) error {
	j := jsonImpression{BidFloorCurrency: "USD"}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	*imp = (Impression)(j)
	return nil
}
