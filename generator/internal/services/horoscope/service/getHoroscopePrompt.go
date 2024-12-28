package service

import (
	"bytes"
	"context"
	"text/template"
	"time"

	"generator/internal/enum/promptCase"
	"generator/internal/services/horoscope/model"
	"generator/internal/services/horoscope/service/utils"
	promptModel "generator/internal/services/promptTemplate/model"
	"pkg/errors"
	"pkg/pointer"
	"pkg/slices"
)

func (s *HoroscopeService) GetHoroscopePrompt(ctx context.Context, req model.GetHoroscopeReq) (res model.GetHoroscopePromptRes, err error) {

	req.DateFrom, req.DateTo = utils.GetDateRangeForTimeframe(req)

	return s.getHoroscopePrompt(ctx, req)
}

func (s *HoroscopeService) getHoroscopePrompt(ctx context.Context, req model.GetHoroscopeReq) (res model.GetHoroscopePromptRes, err error) {

	// Получаем темплейт промпта для генерации гороскопа из базы данных
	promptRes, err := slices.FirstWithError(
		s.promptTemplateService.GetPromptTemplates(ctx, promptModel.GetPromptTemplatesReq{ //nolint:exhaustruct
			Cases: []promptCase.PromptCase{promptCase.CreateHoroscope},
		}),
	)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	// Компилируем го темплейт промпта для генерации гороскопа
	t, err := template.New(string(promptCase.CreateHoroscope)).Parse(promptRes.Template)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	var secondaryZodiacString *string
	if req.SecondaryZodiac != nil {
		secondaryZodiacString = pointer.Pointer(string(*req.SecondaryZodiac))
	}

	// Заполняем темплейт данными
	var buf bytes.Buffer
	if err = t.Execute(&buf, struct {
		DateFrom        time.Time
		DateTo          time.Time
		PrimaryZodiac   string
		SecondaryZodiac *string
		Language        string
		Timeframe       string
		Type            string
	}{
		DateFrom:        req.DateFrom.Time,
		DateTo:          req.DateTo.Time,
		PrimaryZodiac:   string(req.PrimaryZodiac),
		SecondaryZodiac: secondaryZodiacString,
		Language:        string(req.Language),
		Timeframe:       string(req.Timeframe),
		Type:            string(req.HoroscopeType),
	}); err != nil {
		return res, errors.BadRequest.Wrap(err)
	}

	return model.GetHoroscopePromptRes{
		Prompt: buf.String(),
	}, nil
}
