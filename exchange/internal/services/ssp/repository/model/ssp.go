package model

import (
	"fmt"

	"github.com/lib/pq"

	"exchange/internal/config"
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
	"exchange/internal/services/ssp/model"
	"exchange/internal/services/ssp/model/fraudScore"
)

type SSP struct {
	Slug                string                          `db:"slug"`                  // Системное имя
	IsEnable            bool                            `db:"is_enable"`             // Включена ли SSP для входящих запросов
	ClickunderDrumSize  *int32                          `db:"clickunder_drum_size"`  // Размер барабана для кликандер запроса
	ClickunderADMFormat *string                         `db:"clickunder_adm_format"` // Тип ADM ответа кликандера
	Name                string                          `db:"name"`                  // Название
	Timeout             *int32                          `db:"timeout"`               // Таймаут ответа к SSP. В миллисекундах
	IntegrationType     integrationType.IntegrationType `db:"integration_type"`      // Тип интеграции
	EndpointURL         string                          `db:"endpoint_url"`          // URL для интеграции
	SourceTrafficType   pq.StringArray                  `db:"source_traffic_types"`  // Тип трафика по источнику
	BillingType         billingType.BillingType         `db:"billing_type"`          // Способ биллинга SSP
	AuctionSecondPrice  bool                            `db:"auction_second_price"`  // Алгоритм цены. true - второй цены, false - первой цены
	FraudScore          fraudScore.FraudScore           `db:"fraud_score"`           // Выбор проверки FraudScore
	Currency            currency.Currency               `db:"currency"`              // Валюта
	FormatTypes         pq.StringArray                  `db:"format_types"`          // Тип формата
}

func (d *SSP) ConvertToModel() model.SSP {

	formatTypes := make([]formatType.FormatType, 0, len(d.FormatTypes))
	for _, _formatType := range d.FormatTypes {
		formatTypes = append(formatTypes, formatType.FormatType(_formatType))
	}

	sourceTrafficTypes := make([]sourceTrafficType.SourceTrafficType, 0, len(d.SourceTrafficType))
	for _, _sourceTrafficType := range d.SourceTrafficType {
		sourceTrafficTypes = append(sourceTrafficTypes, sourceTrafficType.SourceTrafficType(_sourceTrafficType))
	}

	endpointURL := d.BuildEndpointURL(d.Slug)

	return model.SSP{
		Slug:                d.Slug,
		Name:                d.Name,
		Timeout:             d.Timeout,
		IsEnable:            d.IsEnable,
		IntegrationType:     d.IntegrationType,
		EndpointURL:         endpointURL,
		SourceTrafficTypes:  sourceTrafficTypes,
		BillingType:         d.BillingType,
		AuctionSecondPrice:  d.AuctionSecondPrice,
		Currency:            d.Currency,
		FormatTypes:         formatTypes,
		FraudScore:          d.FraudScore,
		ClickunderADMFormat: d.ClickunderADMFormat,
		ClickunderDrumSize:  d.ClickunderDrumSize,
	}
}

func (d *SSP) BuildEndpointURL(sspSlug string) string {
	return fmt.Sprintf("https://bid.%s/%s", config.Load().Host, sspSlug)
}
