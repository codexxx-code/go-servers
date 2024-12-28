package promptCase

import "pkg/errors"

type PromptCase string

// enums:"createHoroscope"
const (
	CreateHoroscope PromptCase = "createHoroscope"
)

func (z PromptCase) Validate() error {
	switch z {
	case CreateHoroscope:
		return nil
	default:
		return errors.BadRequest.New("Unknown PromptCase", errors.ParamsOption("PromptCase", z))
	}
}
