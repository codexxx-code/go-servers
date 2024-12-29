package beforeResponseToSSP

import (
	"pkg/uuid"
)

type initResponse struct {
	baseLink
}

func (r *initResponse) Apply(dto *beforeResponseToSSP) (err error) {

	// Устанавливаем идентификатор запроса от SSP
	dto.bidResponse.ID = dto.BidRequest.ID

	// Устанавливаем свой идентификатор биддера
	dto.bidResponse.BidID = uuid.New()

	// Устанавливаем ту же валюту, что и валюта запроса
	dto.bidResponse.Currency = dto.RequestCurrency

	return nil
}
