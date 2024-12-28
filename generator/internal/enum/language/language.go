package language

import "pkg/errors"

type Language string

// enums:"russian,english"
const (
	Russian Language = "russian"
	English Language = "english"
)

func (l Language) Validate() error {
	switch l {
	case Russian, English:
		return nil
	default:
		return errors.BadRequest.New("Unknown language", errors.ParamsOption("language", l))
	}
}
