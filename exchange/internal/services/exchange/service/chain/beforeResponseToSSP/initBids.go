package beforeResponseToSSP

import (
	"pkg/openrtb"
	"pkg/uuid"
)

type initBids struct {
	baseLink
}

func (r *initBids) Apply(dto *beforeResponseToSSP) (err error) {

	// Если у нас уже инициализированы биды, то ничего не делаем
	if dto.chainSettings.bidsAlreadyInit {
		return nil
	}

	// Проставляем настройку, что бид инициализирован
	defer func() {
		dto.chainSettings.bidsAlreadyInit = true
	}()

	// Добавляем столько бидов, сколько было импрешенов
	for _, impression := range dto.BidRequest.Impressions {
		dto.bidResponse.SeatBids = append(dto.bidResponse.SeatBids, openrtb.SeatBid{ //nolint:exhaustruct
			Bids: []openrtb.Bid{
				{ //nolint:exhaustruct
					ID:    uuid.New(),    // Генерируем идентификатор бида
					ImpID: impression.ID, // Берем идентификатор импрешена
				},
			},
		})
	}

	return nil
}
