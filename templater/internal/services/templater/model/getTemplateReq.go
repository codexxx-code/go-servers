package model

import "pkg/openrtb"

type GetTemplateReq struct {
	SSPSlug    string             `json:"-" validate:"required"`
	BidRequest openrtb.BidRequest `json:"bidRequest" validate:"required"`
}
