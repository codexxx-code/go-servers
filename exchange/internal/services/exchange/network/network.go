package network

import (
	"github.com/valyala/fasthttp"
)

type ExchangeNetwork struct {
	client *fasthttp.Client
}

func NewExchangeNetwork(client *fasthttp.Client) *ExchangeNetwork {
	return &ExchangeNetwork{client: client}
}
