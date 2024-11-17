package openrtb

import (
	"encoding/json"
	"errors"
)

var (
	ErrInvalidRequestNoID     = errors.New("request has no ID")
	ErrInvalidRequestNoImps   = errors.New("request has no impressions")
	ErrInvalidRequestMultiInv = errors.New("request has multiple inventory sources")
)

// BidRequest is the top-level bid request object contains a globally unique bid request
// or auction ID. This "id" attribute is required as is at least one "imp" (i.e., impression)
// object. Other attributes are optional since an exchange may establish default values.
type BidRequest struct {
	// Unique ID of the bid request, provided by the exchange.
	//
	// Required.
	ID string `json:"id" bson:"id"`

	// Array of Impression objects representing the impressions offered.
	// At least 1 Impression object is required.
	//
	// Required.
	Impressions []Impression `json:"imp" bson:"imp"`

	// Details via a Site object about the publisher’s website. Only applicable and
	// recommended for websites.
	//
	// Recommended.
	Site *Site `json:"site" bson:"site"`

	// Details via an App object about the publisher’s app (i.e., non-browser applications).
	// Only applicable and recommended for apps.
	//
	// Recommended.
	App *App `json:"app" bson:"app"`

	// Details via a Device object about the user’s device to which the impression will be
	// delivered.
	//
	// Recommended.
	Device *Device `json:"device" bson:"device"`

	// Details via a User object about the human user of the device; the advertising audience.
	//
	// Recommended.
	User *User `json:"user" bson:"user"`

	// Indicator of test mode in which auctions are not billable, where:
	//    0 = live mode;
	//    1 = test mode.
	//
	// Default 0.
	Test int `json:"test" bson:"test"`

	// Auction type, where:
	//    1 = First Price;
	//    2 = Second Price Plus.
	// Exchange-specific auction types can be defined using values greater than 500.
	//
	// Default 2.
	AuctionType int `json:"at" bson:"at"`

	// Maximum time in milliseconds the exchange allows for bids to be received including
	// Internet latency to avoid timeout. This value supersedes any a priori guidance from
	// the exchange.
	TimeMax int `json:"tmax" bson:"tmax"`

	// White list of buyer seats (e.g., advertisers, agencies) allowed to bid on this impression.
	// IDs of seats and knowledge of the buyer’s customers to which they refer must be coordinated
	// between bidders and the exchange a priori. At most, only one of wseat and bseat should be
	// used in the same request. Omission of both implies no seat restrictions.
	Seats []string `json:"wseat" bson:"wseat"`

	// Block list of buyer seats (e.g., advertisers, agencies) restricted from bidding on this
	// impression. IDs of seats and knowledge of the buyer’s customers to which they refer must be
	// coordinated between bidders and the exchange a priori. At most, only one of wseat and bseat
	// should be used in the same request. Omission of both implies no seat restrictions.
	BlockedSeats []string `json:"bseat" bson:"bseat"`

	// Flag to indicate if Exchange can verify that the impressions offered represent all of
	// the impressions available in context (e.g., all on the web page, all video spots such
	// as pre/mid/post roll) to support road-blocking:
	//    0 = no or unknown;
	//    1 = yes, the impressions offered represent all that are available.
	//
	// Default 0.
	AllImpressions int `json:"allimps" bson:"allimps"`

	// Array of allowed currencies for bids on this bid request using ISO-4217 alpha codes.
	//
	// Recommended only if the exchange accepts multiple currencies.
	Currencies []string `json:"cur" bson:"cur"`

	// White list of languages for creatives using ISO-639-1-alpha-2. Omission implies
	// no specific restrictions, but buyers would be advised to consider language attribute
	// in the Device and/or Content objects if available.
	Languages []string `json:"wlang" bson:"wlang"`

	// Blocked advertiser categories using the IAB content categories.
	BlockedCategories []ContentCategory `json:"bcat" bson:"bcat"`

	// Block list of advertisers by their domains (e.g., “ford.com”).
	BlockedAdvDomains []string `json:"badv" bson:"badv"`

	// Block list of applications by their platform-specific exchange
	// independent application identifiers. On Android, these should
	// be bundle or package names (e.g., com.foo.mygame).
	// On iOS, these are numeric IDs.
	BlockedApps []string `json:"bapp" bson:"bapp"`

	// A Sorce object that provides data about the inventory source and which entity makes
	// the final decision.
	Source *Source `json:"source" bson:"source"`

	// A Regs object that specifies any industry, legal, or governmental regulations in force
	// for this request.
	Regulations *Regulations `json:"regs" bson:"regs"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}

type jsonBidRequest BidRequest

// UnmarshalJSON custom unmarshaling.
func (r *BidRequest) UnmarshalJSON(data []byte) error {
	alias := jsonBidRequest{AuctionType: 2}
	if err := json.Unmarshal(data, &alias); err != nil {
		return err
	}
	*r = (BidRequest)(alias)
	return nil
}

// Validate the request.
func (r *BidRequest) Validate() error {
	switch {
	case r.ID == "":
		return ErrInvalidRequestNoID
	case len(r.Impressions) == 0:
		return ErrInvalidRequestNoImps
	case r.Site != nil && r.App != nil:
		return ErrInvalidRequestMultiInv
	}

	for i := range r.Impressions {
		imp := r.Impressions[i]
		if err := imp.Validate(); err != nil {
			return err
		}
	}

	return nil
}
