package model

import (
	"exchange/internal/enum/table"
	"exchange/internal/services/dsp/model/dspFilters"
	"exchange/internal/services/dsp/model/dspSort"
)

type FindDSPsReq struct {
	Pagination table.PaginationReq `validate:"required"`
	Sorts      []dspSort.DSPSort
	Filters    dspFilters.DSPFilters
}

func (r FindDSPsReq) Validate() error {
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
