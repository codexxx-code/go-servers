package utils

import (
	dspModel "exchange/internal/services/dsp/model"
	"pkg/errors"
	"pkg/openrtb"
)

func ValidateResponseFromDSP(bidResponse *openrtb.BidResponse, dsp dspModel.DSP) error {

	// Проверяем, есть ли ответ от DSP
	if bidResponse == nil {
		return errors.BadRequest.New("В DTO SendRequestToDSPDTO при пустой ошибке пустая модель bidResponse")
	}

	// Валидируем полученный ответ
	if err := bidResponse.Validate(); err != nil {
		return errors.BadRequest.Wrap(err)
	}

	// Проверяем, что валюта ответа совпадает с валютой DSP
	if string(dsp.Currency) != bidResponse.Currency {
		return errors.BadRequest.New("DSP currency not equal BidResponse currency",
			errors.ParamsOption(
				"DSP slug", dsp.Slug,
				"DSP currency", dsp.Currency,
				"BidResponse currency", bidResponse.Currency,
			),
		)
	}
	return nil
}
