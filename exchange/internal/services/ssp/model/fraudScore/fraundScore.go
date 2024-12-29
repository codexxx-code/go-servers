package fraudScore

import (
	"pkg/errors"
)

type FraudScore string

const (
	Disabled = "disable"   // проверка фраудскором не делается
	RealTime = "real_time" // real-time safeclick
	PreBid   = "pre_bid"   // pre-bid
)

// Validate проверяет, что FraudScore имеет допустимое значение
func (f FraudScore) Validate() error {
	switch f {
	case Disabled, RealTime, PreBid:
		return nil
	default:
		return errors.BadRequest.New("FraudScore undefined")
	}
}
