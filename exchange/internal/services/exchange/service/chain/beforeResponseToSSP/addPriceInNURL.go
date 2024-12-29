package beforeResponseToSSP

import (
	"fmt"

	"pkg/errors"
)

type addPriceInNURL struct {
	baseLink
}

const (
	bidPriceParamField = "bid_price"
)

func (r *addPriceInNURL) Apply(dto *beforeResponseToSSP) error {

	// Проверяем, что NURL уже проставлен
	if !dto.chainSettings.nurlAlreadySet {
		return errors.InternalServer.New("NURL еще не проставлен, а addPriceInNURL уже вызван")
	}

	// TODO: Сделать отдельный флаг в SSP о необходимости добавления цены в NURL и поменять это условие на него
	if !dto.IsClickunder {
		return nil
	}

	// Проходимся по каждому биду из ответа
	for seatBidID, seatBid := range dto.bidResponse.SeatBids {
		for bidID, bid := range seatBid.Bids {

			// Добавляем цену в NURL
			// Некоторые SSP не умеют раскрывать макрос и нам необходимо самим подставить цену
			// Обязательно нужно проверить, что с SSP мы торгуем по первой цене
			dto.bidResponse.SeatBids[seatBidID].Bids[bidID].NoticeURL = fmt.Sprintf("%s&%s=%s", bid.NoticeURL, bidPriceParamField, bid.Price.String()) // TODO: Обыграть по другому
		}
	}

	return nil
}
