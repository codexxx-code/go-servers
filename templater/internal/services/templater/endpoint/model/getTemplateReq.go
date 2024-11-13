package model

import (
	"templater/internal/services/templater/model"

	"pkg/openrtb"
)

type GetTemplateReq struct {
	openrtb.BidRequest
}

func (r GetTemplateReq) ConvertToBusinessModel(sspSlug string) model.GetTemplateReq {
	return model.GetTemplateReq{
		SSPSlug:    sspSlug,
		BidRequest: r.BidRequest,
	}
}
