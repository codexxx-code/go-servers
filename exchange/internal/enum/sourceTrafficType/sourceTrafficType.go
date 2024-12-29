package sourceTrafficType

import (
	"pkg/errors"
)

type SourceTrafficType string

const (
	InApp     = "in_app"
	MobileWeb = "mobile_web"
	Desktop   = "desktop"
	SmartTV   = "smart"
)

func (s SourceTrafficType) Validate() error {
	switch s {
	case InApp, MobileWeb, Desktop, SmartTV:
		return nil
	default:
		return errors.BadRequest.New("SourceTrafficType undefined")
	}
}
