package formatType

import (
	"pkg/errors"
)

type FormatType string

const (
	Banner     = "banner"
	Video      = "video"
	Push       = "push"
	Clickunder = "clickunder"
	Native     = "native"
)

func (f FormatType) Validate() error {
	switch f {
	case Banner, Video, Push, Clickunder, Native:
		return nil
	default:
		return errors.BadRequest.New("FormatType undefined")
	}
}
