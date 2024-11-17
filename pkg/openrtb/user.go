package openrtb

import "encoding/json"

// Human user of the device; audience for advertising.
type User struct {
	// Exchange-specific ID for the user.
	//
	// At least one of id or buyeruid is recommended.
	ID string `json:"id" bson:"id"`

	// Buyer-specific ID for the user as mapped by the exchange for the buyer.
	//
	// At least one of buyeruid or id is recommended.
	BuyerUID string `json:"buyeruid" bson:"buyeruid"`

	// Year of birth as a 4-digit integer.
	YearOfBirth int `json:"yob" bson:"yob"`

	// Gender, where:
	//   M = male;
	//   F = female;
	//   O = known to be other (i.e., omitted is unknown).
	Gender string `json:"gender" bson:"gender"`

	// Comma separated list of keywords, interests, or intent.
	//
	// FIXME: keywords can be a string or an array strings.
	Keywords string `json:"keywords" bson:"keywords"`

	// Optional feature to pass bidder data that was set in the exchange’s cookie.
	// The string must be in base85 cookie safe characters and be in any format.
	// Proper JSON encoding must be used to include “escaped” quotation marks.
	CustomData string `json:"customdata" bson:"customdata"`

	// Location of the user’s home base defined by a Geo object. This is not necessarily
	// their current location.
	Geo *Geo `json:"geo" bson:"geo"`

	// Additional user data. Each Data object represents a different data source.
	Data []Data `json:"data" bson:"data"`

	// Placeholder for exchange-specific extensions to OpenRTB.
	Ext json.RawMessage `json:"ext,omitempty" bson:"ext"`
}
