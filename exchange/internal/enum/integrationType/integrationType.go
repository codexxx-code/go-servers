package integrationType

import (
	"pkg/errors"
)

type IntegrationType string

const (
	ORTB = "ortb"
	VAST = "vast"
	Fedd = "feed"
)

func (t IntegrationType) Validate() error {
	switch t {
	case ORTB, VAST, Fedd:
		return nil
	default:
		return errors.BadRequest.New("IntegrationType undefined")
	}
}
