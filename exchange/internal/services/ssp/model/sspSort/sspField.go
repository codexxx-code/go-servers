package sspSort

import (
	"exchange/internal/services/ssp/repository/sspDDL"
	"pkg/errors"
)

type SSPField int

const (
	Slug = iota + 1
	Name
	Timeout
	IsEnable
	IntegrationType
	SourceTrafficType
	BillingType
	AuctionSecondPrice
	Currency
	FormatTypes
	ClickunderDrumSize
	ClickunderAdmFormat
	FraudScore
)

var mappingModelToDDL = map[SSPField]string{
	Slug:                sspDDL.ColumnSlug,
	Name:                sspDDL.ColumnName,
	Timeout:             sspDDL.ColumnTimeout,
	IsEnable:            sspDDL.ColumnIsEnable,
	IntegrationType:     sspDDL.ColumnIntegrationType,
	SourceTrafficType:   sspDDL.ColumnSourceTrafficTypes,
	BillingType:         sspDDL.ColumnBillingType,
	AuctionSecondPrice:  sspDDL.ColumnAuctionSecondPrice,
	Currency:            sspDDL.ColumnCurrency,
	FormatTypes:         sspDDL.ColumnFormatTypes,
	ClickunderDrumSize:  sspDDL.ColumnClickunderDrumSize,
	ClickunderAdmFormat: sspDDL.ColumnClickunderADMFormat,
	FraudScore:          sspDDL.ColumnFraudScore,
}

// Validate проверяет, что SSPField имеет допустимое значение
func (s SSPField) Validate() error {
	switch s {
	case
		Slug,
		Name,
		Timeout,
		IsEnable,
		IntegrationType,
		SourceTrafficType,
		BillingType,
		AuctionSecondPrice,
		Currency,
		FormatTypes,
		ClickunderDrumSize,
		ClickunderAdmFormat,
		FraudScore:

		return nil
	default:
		return errors.BadRequest.New("SSPField undefined")
	}
}

func (s SSPField) ConvertToDDL() string {
	return mappingModelToDDL[s]
}
