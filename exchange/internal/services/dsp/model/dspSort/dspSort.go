package dspSort

import "exchange/internal/enum/table/direction"

type DSPSort struct {
	DSPField  DSPSortField        `validate:"required"` // название поля
	Direction direction.Direction `validate:"required"` // тип сортировки ASC, DESC
}

func (s DSPSort) Validate() error {
	if err := s.DSPField.Validate(); err != nil {
		return err
	}

	if err := s.Direction.Validate(); err != nil {
		return err
	}

	return nil
}
