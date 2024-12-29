package billingType

import (
	"pkg/errors"
)

type BillingType string

const (
	NURL        = "nurl"
	BURL        = "burl"
	ADM         = "adm"
	Imp         = "imp"
	NURLAndBURL = "nurl_and_burl"
)

// Validate проверяет, что BillingType имеет допустимое значение
func (b BillingType) Validate() error {
	switch b {
	case NURL, BURL, ADM, Imp, NURLAndBURL:
		return nil
	default:
		return errors.BadRequest.New("BillingType undefined")
	}
}
