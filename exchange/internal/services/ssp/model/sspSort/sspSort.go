package sspSort

import (
	"exchange/internal/enum/table/direction"
)

type SSPSort struct {
	SSPField  SSPField            `validate:"required"` // название поля
	Direction direction.Direction `validate:"required"` // тип сортировки ASC, DESC
}

func (s SSPSort) Validate() error {
	if err := s.SSPField.Validate(); err != nil {
		return err
	}

	if err := s.Direction.Validate(); err != nil {
		return err
	}

	return nil
}
