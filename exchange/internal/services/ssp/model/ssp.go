package model

import (
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
	"exchange/internal/services/ssp/model/fraudScore"
)

type SSP struct {
	Slug                string                                // Системное имя
	IsEnable            bool                                  // Включена ли SSP для входящих запросов
	ClickunderDrumSize  *int32                                // Размер барабана для кликандер запроса
	ClickunderADMFormat *string                               // Тип ADM ответа кликандера
	Name                string                                // Название
	Timeout             *int32                                // Таймаут ответа к SSP. В миллисекундах
	IntegrationType     integrationType.IntegrationType       // Тип интеграции
	EndpointURL         string                                // URL для интеграции
	SourceTrafficTypes  []sourceTrafficType.SourceTrafficType // Тип трафика по источнику
	BillingType         billingType.BillingType               // Способ биллинга SSP
	AuctionSecondPrice  bool                                  // Алгоритм цены. true - второй цены, false - первой цены
	FraudScore          fraudScore.FraudScore                 // Выбор проверки FraudScore
	Currency            currency.Currency                     // Валюта
	FormatTypes         []formatType.FormatType
	DSPs                []string // Модифиакторы связки SSP и DSP
}
