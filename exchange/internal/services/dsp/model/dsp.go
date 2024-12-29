package model

import (
	"exchange/internal/enum/billingType"
	"exchange/internal/enum/currency"
	"exchange/internal/enum/formatType"
	"exchange/internal/enum/integrationType"
	"exchange/internal/enum/sourceTrafficType"
)

type DSP struct {
	Name                     string                                // Человекочитаемое название
	Slug                     string                                // Системное имя
	URL                      string                                // URL запроса
	Currency                 currency.Currency                     // Код валюты по ISO-4217
	AuctionSecondPrice       bool                                  // Показатель, аукцион второй или первой цены у DSP
	BillingURLType           billingType.BillingType               // Настройка, по какому URL биллим DSP
	IsSupportMultiimpression bool                                  // Показатель, поддерживает ли DSP мультиимпрешены
	IsEnable                 bool                                  // Показатель, показывающий статус DSP
	IntegrationType          integrationType.IntegrationType       // Тип интеграции
	SourceTrafficTypes       []sourceTrafficType.SourceTrafficType // Тип источника трафика
	FormatTypes              []formatType.FormatType               // Тип формата
}
