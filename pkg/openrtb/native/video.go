package native

import (
	"pkg/openrtb"
)

// VideoRequest is an video element in Native ad.
//
// This corresponds to the VideoRequest object of OpenRTB.
//
// Exchange implementers can impose their own specific restrictions.
// Here are the required attributes of the VideoRequest Object. For optional
// attributes please refer to OpenRTB.
type VideoRequest = openrtb.Video

// VideoResponse is an video element in Native ad.
//
// Corresponds to the Video Object in the request, yet containing a value of
// a conforming VAST tag as a value.
type VideoResponse struct {
	VAST string `json:"vasttag"` // VAST XML.
}
