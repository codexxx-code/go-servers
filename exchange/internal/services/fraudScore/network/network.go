package network

import (
	"github.com/valyala/fasthttp"
)

// FraudScoreNetwork представляет HTTP-клиент для взаимодействия с Fraudscore API.
type FraudScoreNetwork struct {
	client *fasthttp.Client
	url    string
	key    string
}

// NewFraudScoreNetwork создает новый клиент FraudScoreNetwork.
func NewFraudScoreNetwork(clientHTTP *fasthttp.Client, key string) *FraudScoreNetwork {
	return &FraudScoreNetwork{
		client: clientHTTP,
		url:    "https://check.fraudscore.mobi",
		key:    key,
	}
}
