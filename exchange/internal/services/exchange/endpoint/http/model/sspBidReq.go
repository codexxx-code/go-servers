package model

import (
	"exchange/internal/services/exchange/model"
	"pkg/openrtb"
)

type SSPBidReq struct {
	openrtb.BidRequest
}

func (s *SSPBidReq) ConvertToBusinessModel(sspSlug string) model.SSPBidReq {
	return model.SSPBidReq{
		BidRequest: s.BidRequest,
		SSPSlug:    sspSlug,
	}
}
