package beforeResponseToSSP

import (
	"pkg/currencyConverter"
	"pkg/decimal"
	"pkg/errors"
)

type checkMaxPriceClickunder struct {
	baseLink
}

const maxPriceInRUBClickunder = 1000

func (r *checkMaxPriceClickunder) Apply(dto *beforeResponseToSSP) (err error) {

	// Если запрос не кликандер, для обычных ответов есть другой чейн
	if !dto.IsClickunder {
		return nil
	}

	// Проходимся по каждой ставке
	for _, seatBid := range dto.bidResponse.SeatBids {
		for _, bid := range seatBid.Bids {

			// Если ставка больше 1000 рублей или эквивалента
			priceInRUB, err := currencyConverter.Convert(
				bid.Price,
				dto.bidResponse.Currency,
				"RUB",
				dto.CurrencyRates,
			)
			if err != nil {
				return err
			}
			if priceInRUB.GreaterThanOrEqual(decimal.NewFromInt(maxPriceInRUBClickunder)) {
				return errors.BadRequest.Wrap(ErrPriceIsTooHigh,
					errors.ParamsOption(
						"price", bid.Price,
						"currency", dto.bidResponse.Currency,
						"requestID", dto.ExchangeID,
					),
				)
			}
		}
	}
	return nil
}
