package model

import (
	"time"

	"pkg/decimal"
	"pkg/openrtb"
)

type DSPResponse struct {
	ExchangeID string `bson:"requestID"` // Идентификатор запроса, формируется нами

	ExchangeBidID string              `bson:"_id"`          // Уникальный идентификатор записи/ставки
	BidResponse   openrtb.BidResponse `bson:"bidResponse"`  // Изначальный ответ DSP
	SeatBidIndex  int                 `bson:"seatBidIndex"` // Индекс победившего SeatBid в ответе DSP
	BidIndex      int                 `bson:"bidIndex"`     // Индекс победившего Bid в ответе DSP

	// Цена, по которой будем биллить DSP, она в валюте DSP и уже с учетом
	// нашей маржи, так как мы ее добавили к bidFloor
	BillingPriceInDSPCurrency decimal.Decimal `json:"price" bson:"billingPriceInDSPCurrency"`

	SlugDSP         string    `bson:"slugDSP"`         // Строковый идентификатор победившей DSP
	RecordDateTime  time.Time `bson:"dateTime"`        // Дата и время создания записи
	OriginRequestID string    `bson:"originRequestID"` // Идентификатор bidRequest.id изначального запроса

	// Коэффициент конвертации валюты SSP в базовую валюту
	// Когда умножаем полученную от SSP цену на этот коэффициент, получаем цену в базовой валюте
	// priceFromSSP * currencySSPCoefficient = priceInBaseCurrency
	// Это поле существует для того, чтобы мы биллили SSP по тому же курсу, что и был на аукционе
	CurrencySSPCoefficient decimal.Decimal `bson:"currencySSPCoefficient"`
	SSPCurrency            string          `bson:"sspCurrency"` // Валюта SSP запроса
	SlugSSP                string          `bson:"slugSSP"`     // Строковый идентификатор ssp

	// Коэффициент конвертации валюты DSP в базовую валюту
	// Когда умножаем цену, по которой мы должны биллить DSP на этот коэффициент, получаем цену в базовой валюте
	// billingPriceForDSP * currencyDSPCoefficient = priceInBaseCurrency
	// Это поле существует для того, чтобы мы биллили DSP по тому же курсу, что и был на аукционе
	CurrencyDSPCoefficient decimal.Decimal `bson:"currencyDSPCoefficient"` // Коэффициент конвертации валюты DSP в базовую валюту
	DSPCurrency            string          `bson:"dspCurrency"`            // Валюта DSP ответа

	RequestPublisherID string `bson:"requestPublisherID"` // Идентификатор паблишера на запросе

	DrumSize *int32 // Размер барабана
}
