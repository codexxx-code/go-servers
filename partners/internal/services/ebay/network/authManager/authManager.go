package authManager

import (
	"context"
	"encoding/base64"
	"fmt"
	"sync"
	"time"

	"partners/internal/config"
	"partners/internal/services/ebay/network/model"
)

type bearer struct {
	mu                     *sync.Mutex
	bearerTokenExpiredDate time.Time
	bearerToken            string
}

type EbayAuthManager struct {
	ebayNetwork EbayNetwork

	basicToken string

	bearer bearer
}

// Из-за циклической ссылки не скомпилится код
// var _ EbayNetwork = new(ebayNetwork.EbayNetwork)

type EbayNetwork interface {
	Auth(context.Context, model.AuthReq) (model.AuthRes, error)
}

func NewEbayAuthManager(
	config config.EbayConfig,
	ebayNetwork EbayNetwork,
) *EbayAuthManager {

	// Кодируем в base64 строку вида "client_id:client_secret"
	basicToken := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", config.ClientID, config.ClientSecret)))

	return &EbayAuthManager{
		ebayNetwork: ebayNetwork,
		basicToken:  basicToken,
		bearer: bearer{
			bearerTokenExpiredDate: time.Time{},
			bearerToken:            "",
			mu:                     &sync.Mutex{},
		},
	}
}
