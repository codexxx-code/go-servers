package beforeResponseToSSP

import (
	"pkg/currencyConverter"
	"pkg/decimal"
	"pkg/errors"
)

type checkMaxPrice struct {
	baseLink
}

const maxPriceInRUB = 500

var ErrPriceIsTooHigh = errors.New("BillingPriceInDSPCurrency is too high")

func (r *checkMaxPrice) Apply(dto *beforeResponseToSSP) (err error) {

	// Если запрос кликандер, то ничего не делаем, потому что проверка есть в другом чейне
	if dto.IsClickunder {
		return nil
	}

	// Проходимся по каждой ставке
	for _, seatBid := range dto.bidResponse.SeatBids {
		for _, bid := range seatBid.Bids {

			// Если ставка больше 500 рублей или эквивалента
			priceInRUB, err := currencyConverter.Convert(
				bid.Price,
				dto.bidResponse.Currency,
				"RUB",
				dto.CurrencyRates,
			)
			if err != nil {
				return err
			}
			if priceInRUB.GreaterThanOrEqual(decimal.NewFromInt(maxPriceInRUB)) {
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
