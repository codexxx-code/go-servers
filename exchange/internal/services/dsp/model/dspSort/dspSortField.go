package dspSort

import (
	"exchange/internal/services/dsp/repository/dspDDL"
	"pkg/errors"
)

type DSPSortField int

const (
	Slug = iota + 1
	NAME
	EndpointUrl
	Currency
	AuctionSecondPrice
	BillingType
	IsEnable
	IntegrationType
	SourceTrafficType
	FormatTypes
	IsSupportMultiimpression
)

var mappingModelToDDL = map[DSPSortField]string{
	Slug:                     dspDDL.ColumnSlug,
	NAME:                     dspDDL.ColumnName,
	EndpointUrl:              dspDDL.ColumnUrl,
	Currency:                 dspDDL.ColumnCurrency,
	AuctionSecondPrice:       dspDDL.ColumnAuctionSecondPrice,
	BillingType:              dspDDL.ColumnBillingUrlType,
	IsEnable:                 dspDDL.ColumnIsEnable,
	IntegrationType:          dspDDL.ColumnIntegrationType,
	SourceTrafficType:        dspDDL.ColumnSourceTrafficTypes,
	FormatTypes:              dspDDL.ColumnFormatTypes,
	IsSupportMultiimpression: dspDDL.ColumnIsSupportMultiimpression,
}

func (s DSPSortField) ConvertToDDL() string {
	return mappingModelToDDL[s]
}

// Validate проверяет, что DSPSortField имеет допустимое значение
func (s DSPSortField) Validate() error {
	switch s {
	case
		Slug,
		NAME,
		EndpointUrl,
		Currency,
		AuctionSecondPrice,
		BillingType,
		IsEnable,
		IntegrationType,
		SourceTrafficType,
		FormatTypes,
		IsSupportMultiimpression:

		return nil
	default:
		return errors.BadRequest.New("DSPSortField undefined")
	}
}
