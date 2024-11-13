package service

import (
	"context"
	"encoding/json"

	"pkg/errors"
	"pkg/openrtb"
	"templater/internal/services/templater/model"
)

func (s *TemplaterService) GetTemplate(ctx context.Context, req model.GetTemplateReq) (res openrtb.BidResponse, err error) {

	// Получаем теплейт по слагу SSP
	getTemplateRes, err := s.templaterRepository.GetTemplate(ctx, req)
	if err != nil {
		return res, err
	}

	// Парсим темплейт в bidResponse
	if err = json.Unmarshal([]byte(getTemplateRes.Template), &res); err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	return res, nil
}
