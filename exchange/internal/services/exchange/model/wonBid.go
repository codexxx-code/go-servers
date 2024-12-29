package model

import (
	"pkg/decimal"
)

type WonBid struct {
	BidPointer
	BillingPriceInDSPCurrency decimal.Decimal // Цена, по которой мы будем биллить DSP в валюте DSP
}
