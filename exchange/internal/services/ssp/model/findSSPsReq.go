package model

import (
	"exchange/internal/enum/table"
	"exchange/internal/services/ssp/model/sspFilters"
	"exchange/internal/services/ssp/model/sspSort"
)

type FindSSPsReq struct {
	Pagination table.PaginationReq `validate:"required"`
	Sorts      []sspSort.SSPSort
	Filters    sspFilters.SSPFilters
}

func (r FindSSPsReq) Validate() error {
	for _, sort := range r.Sorts {
		if err := sort.Validate(); err != nil {
			return err
		}
	}

	if err := r.Filters.Validate(); err != nil {
		return err
	}

	return nil
}
