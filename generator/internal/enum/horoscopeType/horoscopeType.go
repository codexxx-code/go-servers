package horoscopeType

import "pkg/errors"

type HoroscopeType string

// enums:"single,couple"
const (
	Single HoroscopeType = "single"
	Couple HoroscopeType = "couple"
)

func (h HoroscopeType) Validate() error {
	switch h {
	case Single, Couple:
		return nil
	default:
		return errors.BadRequest.New("Unknown HoroscopeType", errors.ParamsOption("HoroscopeType", h))
	}
}
