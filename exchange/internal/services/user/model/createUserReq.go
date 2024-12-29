package model

import (
	"exchange/internal/enum/permission"
)

type CreateUserReq struct {
	LastName      string                  `validate:"min=1,max=256"`
	FirstName     string                  `validate:"min=1,max=256"`
	Email         string                  `validate:"required,min=1,max=256,email"`
	Permissions   []permission.Permission `validate:"eq=1|eq=2|eq=3|eq=4|eq=5|eq=6|eq=7"`
	AuthorID      *string                 `validate:"min=1,max=1000"`
	Password      string                  `validate:"required,min=4,max=255"`
	RetryPassword string                  `validate:"required,min=4,max=255"`
}

func (c CreateUserReq) Validate() error {
	for _, permission := range c.Permissions {
		if err := permission.Validate(); err != nil {
			return err
		}
	}

	return nil
}
