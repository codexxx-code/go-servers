package openrtb

import (
	"encoding/json"
	"errors"
)

// Validation errors.
var (
	ErrInvalidSeatBidNoBids = errors.New("seatbid has no bids")
)

// SeatBid contains seat information. At least one of Bid is required.
// A bid response can contain multiple "seatbid” objects, each on behalf of a different bidder seat.
// SeatBid object can contain multiple bids each pertaining to a different impression on behalf of a seat.
// Each "bid” object must include the impression ID to which it pertains as well as the bid price.
// Group attribute can be used to specify if a seat is willing to accept any impressions that it can win
// (default) or if it is only interested in winning any if it can win them all (i.e., all or nothing).
type SeatBid struct {
	// Array of 1+ Bid objects each related to an impression. Multiple bids can relate
	// to the same impression.
	Bids []Bid `json:"bid" bson:"bid"`

	// ID of the buyer seat (e.g., advertiser, agency) on whose behalf this bid is made.
	Seat string `json:"seat" bson:"seat"`

	//    0 = impressions can be won individually;
	//    1 = impressions must be won or lost as a group.
	//
	// Default 0.
	Group int `json:"group" bson:"group"`

	// Placeholder for bidder-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}

// Validate required attributes.
func (sb *SeatBid) Validate() error {
	if len(sb.Bids) == 0 {
		return ErrInvalidSeatBidNoBids
	}

	for i := range sb.Bids {
		bid := sb.Bids[i]
		if err := bid.Validate(); err != nil {
			return err
		}
	}

	return nil
}
