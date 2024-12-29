package idType

import "pkg/errors"

// IDType - Возможные типы передаваемых идентификаторов в запросе
type IDType string

const (
	ExchangeID    IDType = "request"
	ExchangeBidID IDType = "bid"
)

func (u IDType) Validate() error {
	switch u {
	case ExchangeID, ExchangeBidID:
		return nil
	default:
		return errors.BadRequest.New("ID type not declared in our system")
	}
}
