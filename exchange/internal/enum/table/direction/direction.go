package direction

import (
	"pkg/errors"
)

type Direction int

const (
	Asc = iota + 1
	Desc
)

// Validate проверяет, что Direction имеет допустимое значение
func (d Direction) Validate() error {
	switch d {
	case
		Asc,
		Desc:
		return nil
	default:
		return errors.BadRequest.New("Direction undefined")
	}
}
