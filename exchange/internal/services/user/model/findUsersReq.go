package model

import (
	"exchange/internal/enum/table"
	"exchange/internal/services/user/model/userFilters"
	"exchange/internal/services/user/model/userSort"
)

type FindUsersReq struct {
	Pagination table.PaginationReq `validate:"required"`
	Sorts      []userSort.UserSort
	Filters    userFilters.UserFilters
}

func (f FindUsersReq) Validate() error {

	for _, sort := range f.Sorts {
		if err := sort.Validate(); err != nil {
			return err
		}
	}

	return f.Filters.Validate()
}
