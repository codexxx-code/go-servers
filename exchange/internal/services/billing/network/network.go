package network

import (
	"github.com/valyala/fasthttp"
)

type BillingNetwork struct {
	client *fasthttp.Client
}

func NewBillingNetwork(client *fasthttp.Client) *BillingNetwork {
	return &BillingNetwork{client: client}
}
