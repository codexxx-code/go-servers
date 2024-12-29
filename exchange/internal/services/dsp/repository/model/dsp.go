package model

import (
	"github.com/lib/pq"

	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
	"exchange/internal/services/dsp/model"
)

type DSP struct {
	Name                     string                          `db:"name"`                       // Человекочитаемое название
	Slug                     string                          `db:"slug"`                       // Системное имя
	URL                      string                          `db:"url"`                        // URL запроса
	Currency                 currency.Currency               `db:"currency"`                   // Код валюты по ISO-4217
	AuctionSecondPrice       bool                            `db:"auction_second_price"`       // Показатель, аукцион второй или первой цены у DSP
	BillingURLType           billingType.BillingType         `db:"billing_url_type"`           // Настройка, по какому URL биллим DSP
	IsSupportMultiimpression bool                            `db:"is_support_multiimpression"` // Показатель, поддерживает ли DSP мультиимпрешены
	IsEnable                 bool                            `db:"is_enable"`                  // Показатель, показывающий статус DSP
	IntegrationType          integrationType.IntegrationType `db:"integration_type"`           // Тип интеграции
	SourceTrafficType        pq.StringArray                  `db:"source_traffic_types"`       // Тип источника трафика
	FormatTypes              pq.StringArray                  `db:"format_types"`               // Тип формата
}

func (d *DSP) ConvertToModel() model.DSP {

	formatTypes := make([]formatType.FormatType, 0, len(d.FormatTypes))
	for _, _formatType := range d.FormatTypes {
		formatTypes = append(formatTypes, formatType.FormatType(_formatType))
	}

	sourceTrafficTypes := make([]sourceTrafficType.SourceTrafficType, 0, len(d.SourceTrafficType))
	for _, _sourceTrafficType := range d.SourceTrafficType {
		sourceTrafficTypes = append(sourceTrafficTypes, sourceTrafficType.SourceTrafficType(_sourceTrafficType))
	}

	return model.DSP{
		Name:                     d.Name,
		Slug:                     d.Slug,
		URL:                      d.URL,
		Currency:                 d.Currency,
		AuctionSecondPrice:       d.AuctionSecondPrice,
		BillingURLType:           d.BillingURLType,
		IsSupportMultiimpression: d.IsSupportMultiimpression,
		IsEnable:                 d.IsEnable,
		IntegrationType:          d.IntegrationType,
		SourceTrafficTypes:       sourceTrafficTypes,
		FormatTypes:              formatTypes,
	}
}
