package userSort

import (
	"exchange/internal/enum/table/direction"
)

type UserSort struct {
	UserField UserSortField       `validate:"required"` // название поля
	Direction direction.Direction `validate:"required"` // тип сортировки ASC, DESC
}

func (u UserSort) Validate() error {
	if err := u.UserField.Validate(); err != nil {
		return err
	}
	if err := u.Direction.Validate(); err != nil {
		return err
	}
	return nil
}
