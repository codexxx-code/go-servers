package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"partners/internal/services/ebay/network/model"
	"pkg/errors"
)

const authPath = "/identity/v1/oauth2/token"

func (n *EbayNetwork) Auth(ctx context.Context, req model.AuthReq) (res model.AuthRes, err error) {

	// Формируем тело запроса
	body := url.Values{
		"grant_type": []string{"client_credentials"},
		"scope":      []string{"https://api.ebay.com/oauth/api_scope"},
	}.Encode()

	// Убираем экранирование символов
	body, err = url.QueryUnescape(body)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	// Формируем заголовки запроса
	headers := map[string]string{
		"Content-Type":  "application/x-www-form-urlencoded",
		"Authorization": fmt.Sprintf("Basic %s", req.BasicToken),
	}

	// Отправляем запрос
	resp, err := n.sendRequest(sendRequestDTO{
		ctx:       ctx,
		method:    http.MethodPost,
		path:      authPath,
		urlValues: nil,
		headers:   headers,
		withAuth:  false,
		body:      []byte(body),
	})
	if err != nil {
		return res, err
	}
	defer resp.Body.Close()

	// Читаем ответ
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	return res, nil
}
