package model

import (
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
)

type UpdateDSPReq struct {
	Slug                     string            `validate:"required"`
	Name                     string            `validate:"required"`
	EndpointURL              string            `validate:"required"`
	Currency                 currency.Currency `validate:"required"`
	AuctionSecondPrice       bool
	BillingType              billingType.BillingType `validate:"required"`
	IsEnable                 bool
	IntegrationType          integrationType.IntegrationType       `validate:"required"`
	SourceTrafficTypes       []sourceTrafficType.SourceTrafficType `validate:"required"`
	FormatTypes              []formatType.FormatType               `validate:"required"`
	IsSupportMultiimpression bool
}

func (r UpdateDSPReq) Validate() error {

	if err := r.Currency.Validate(); err != nil {
		return err
	}

	if err := r.BillingType.Validate(); err != nil {
		return err
	}

	if err := r.IntegrationType.Validate(); err != nil {
		return err
	}

	for _, sourceTrafficType := range r.SourceTrafficTypes {
		if err := sourceTrafficType.Validate(); err != nil {
			return err
		}
	}

	for _, formatType := range r.FormatTypes {
		if err := formatType.Validate(); err != nil {
			return err
		}
	}

	return nil
}
