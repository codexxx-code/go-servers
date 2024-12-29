package beforeResponseToSSP

import (
	"pkg/uuid"

	"pkg/openrtb"
)

type initBidsForClikunder struct {
	baseLink
}

func (r *initBidsForClikunder) Apply(dto *beforeResponseToSSP) (err error) {

	// Если у нас уже инициализированы биды, то ничего не делаем
	if dto.chainSettings.bidsAlreadyInit {
		return nil
	}

	// Если у нас запрос не на барабан, то ничего не делаем
	if !dto.IsClickunder {
		return nil
	}

	// Проставляем настройку, что бид инициализирован
	defer func() {
		dto.chainSettings.bidsAlreadyInit = true
	}()

	// Для барабана у нас может быть только один бид
	dto.bidResponse.SeatBids = []openrtb.SeatBid{
		{ //nolint:exhaustruct
			Bids: []openrtb.Bid{
				{ //nolint:exhaustruct
					ID:    uuid.New(),                       // Генерируем идентификатор бида
					ImpID: dto.BidRequest.Impressions[0].ID, // Берем валюту первого импрешена
				},
			},
		},
	}

	return nil
}
