package dspFilters

import (
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
)

type DSPFilters struct {
	Slugs                    []string                              // Список уникальных идентификаторов DSP для фильтрации
	Names                    []string                              // Список имен DSP для фильтрации
	EndpointURLs             []string                              // Список endpoint URL для фильтрации
	Currencies               []currency.Currency                   // Список валют для фильтрации
	AuctionSecondPrice       *bool                                 // Флаг, указывающий, используется ли аукцион второго предложения
	BillingTypes             []billingType.BillingType             // Список типов биллинга для фильтрации
	IsEnable                 *bool                                 // Флаг активности DSP
	IntegrationTypes         []integrationType.IntegrationType     // Список типов интеграции для фильтрации
	SourceTrafficTypes       []sourceTrafficType.SourceTrafficType // Список типов исходного трафика для фильтрации
	FormatTypes              []formatType.FormatType               // Список типов форматов для фильтрации
	IsSupportMultiimpression *bool
}

func (s DSPFilters) Validate() error {
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

	return nil
}
