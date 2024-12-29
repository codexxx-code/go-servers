package sspFilters

import (
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
	"exchange/internal/services/ssp/model/fraudScore"
)

type SSPFilters struct {
	Slugs               []string                              // Список slugов для фильтрации
	Names               []string                              // Список имен для фильтрации
	IsEnable            *bool                                 // Фильтр по состоянию включения
	IntegrationTypes    []integrationType.IntegrationType     // Фильтрация по типам интеграции
	SourceTrafficTypes  []sourceTrafficType.SourceTrafficType // Фильтрация по типам трафика
	BillingTypes        []billingType.BillingType             // Фильтрация по типам биллинга
	AuctionSecondPrice  *bool                                 // Фильтрация по алгоритму аукциона
	Currencies          []currency.Currency                   // Фильтрация по валютам
	FormatTypes         []formatType.FormatType               // Фильтрация по типам форматов
	ClickunderDrumSize  *int32                                // Фильтрация по размеру барабана для кликандер
	ClickunderADMFormat *string                               // Фильтрация по формату с макросом
	FraudScores         []fraudScore.FraudScore               // Фильтрация по типам фраудскора
}

func (s SSPFilters) Validate() error {
	for _, currency := range s.Currencies {
		if err := currency.Validate(); err != nil {
			return err
		}
	}

	for _, billingType := range s.BillingTypes {
		if err := billingType.Validate(); err != nil {
			return err
		}
	}

	for _, integrationType := range s.IntegrationTypes {
		if err := integrationType.Validate(); err != nil {
			return err
		}
	}

	for _, sourceTrafficType := range s.SourceTrafficTypes {
		if err := sourceTrafficType.Validate(); err != nil {
			return err
		}
	}

	for _, formatType := range s.FormatTypes {
		if err := formatType.Validate(); err != nil {
			return err
		}
	}

	for _, fraudScore := range s.FraudScores {
		if err := fraudScore.Validate(); err != nil {
			return err
		}
	}

	return nil
}
