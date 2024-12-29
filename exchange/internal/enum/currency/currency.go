package currency

import (
	"pkg/errors"
)

type Currency string

const (
	RUB = "RUB"
	USD = "USD"
)

// Validate проверяет, что Currency имеет допустимое значение
func (c Currency) Validate() error {
	switch c {
	case RUB, USD:
		return nil
	default:
		return errors.BadRequest.New("Currency undefined")
	}
}
