package model

import (
	"exchange/internal/enum/billingType"
	"exchange/internal/services/exchange/model/idType"
)

type BillURLReq struct {
	ID             string                  `schema:"-" validate:"required"`
	MacrosPrice    string                  `schema:"price"`
	HardcodedPrice string                  `schema:"bid_price"`
	BillingType    billingType.BillingType `schema:"url_type" validate:"required"`
	IDType         idType.IDType           `schema:"id_type" validate:"required"`
	SSPSlug        string                  `schema:"ssp_slug"`
}

func (s BillURLReq) Validate() error {
	if err := s.BillingType.Validate(); err != nil {
		return err
	}
	if err := s.IDType.Validate(); err != nil {
		return err
	}

	return nil
}
