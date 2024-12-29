package model

import (
	"pkg/openrtb"
)

type SSPBidReq struct {
	BidRequest openrtb.BidRequest
	SSPSlug    string
}
