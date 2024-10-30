package network

import (
	"context"
	"net/http"

	"partners/internal/config"
	"partners/internal/services/ebay/network/authManager"
)

type EbayNetwork struct {
	httpClient  http.Client
	baseHost    string
	authManager AuthManager
}

var _ AuthManager = new(authManager.EbayAuthManager)

type AuthManager interface {
	GetToken(ctx context.Context) (string, error)
}

func NewEbayNetwork(config config.EbayConfig) *EbayNetwork {

	ebayNetwork := &EbayNetwork{
		authManager: nil,
		httpClient: http.Client{
			Transport:     nil,
			CheckRedirect: nil,
			Jar:           nil,
			Timeout:       0,
		},
		baseHost: baseHostMap[config.IsSandbox],
	}

	// Делаем запрос на авторизацию
	authManager := authManager.NewEbayAuthManager(config, ebayNetwork)

	// Передаем менеджер авторизации в сеть
	ebayNetwork.authManager = authManager

	return ebayNetwork
}

var baseHostMap = map[bool]string{
	true:  "https://api.sandbox.ebay.com",
	false: "https://api.ebay.com",
}
