package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var (
	ErrInvalidDealNoID = errors.New("deal has no ID")
)

// Collection of private marketplace (PMP) deals applicable to this impression.
type PMP struct {
	// Indicator of auction eligibility to seats named in the Direct Deals object, where:
	//    0 = all bids are accepted;
	//    1 = bids are restricted to the deals specified and the terms thereof.
	//
	// Default 0.
	Private int `json:"private_auction" bson:"private_auction"`

	// Array of Deal objects that convey the specific deals applicable to this impression.
	Deals []Deal `json:"deals" bson:"deals"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}

// Deal terms pertaining to this impression between a seller and buyer.
type Deal struct {
	// A unique identifier for the direct deal.
	//
	// Required.
	ID string `json:"id" bson:"id"`

	// Minimum bid for this impression expressed in CPM.
	//
	// Default 0.
	BidFloor float64 `json:"bidfloor" bson:"bidfloor"`

	// Currency specified using ISO-4217 alpha codes. This may be different from bid currency
	// returned by bidder if this is allowed by the exchange.
	//
	// Default USD.
	BidFloorCurrency string `json:"bidfloorcur" bson:"bidfloorcur"`

	// Optional override of the overall auction type of the bid request, where:
	//    1 = First Price;
	//    2 = Second Price Plus;
	//    3 = the value passed in bidfloor is the agreed upon deal price.
	// Additional auction types can be defined by the exchange.
	AuctionType int `json:"at" bson:"at"`

	// Whitelist of buyer seats (e.g., advertisers, agencies) allowed to bid on this deal.
	// IDs of seats and the buyer’s customers to which they refer must be coordinated between
	// bidders and the exchange a priori. Omission implies no seat restrictions.
	Seats []string `json:"wseat" bson:"wseat"`

	// Array of advertiser domains (e.g., advertiser.com) allowed to bid on this deal.
	// Omission implies no advertiser restrictions.
	AdvDomains []string `json:"wadomain" bson:"wadomain"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}

// Validate the PMP object.
func (p *PMP) Validate() error {
	for i := range p.Deals {
		deal := p.Deals[i]
		if err := deal.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// Validate the Deal object.
func (d *Deal) Validate() error {
	if d.ID == "" {
		return ErrInvalidDealNoID
	}
	return nil
}

type jsonDeal Deal

// UnmarshalJSON custom unmarshalling.
func (d *Deal) UnmarshalJSON(data []byte) error {
	j := jsonDeal{BidFloorCurrency: "USD"}
	if err := json.Unmarshal(data, &j); err != nil {
		return err
	}
	*d = (Deal)(j)
	return nil
}
