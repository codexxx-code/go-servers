package network

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"partners/internal/services/ebay/network/model"
	"pkg/errors"
)

var baseDomainMap = map[bool]string{
	true:  "sandbox.ebay.com",
	false: "ebay.com",
}

func getPath(isSandbox bool, subdomain, path string) string {
	return fmt.Sprintf("https://%s.%s%s", subdomain, baseDomainMap[isSandbox], path)
}

const (
	authPath      = "/identity/v1/oauth2/token"
	authSubdomain = "api"
)

func (s *EbayNetwork) Auth(ctx context.Context, req model.AuthReq) (res model.AuthRes, err error) {

	// Получаем путь для запроса на авторизацию
	authPath := getPath(s.isSandbox, authSubdomain, authPath)

	// Формируем тело запроса
	body := url.Values{
		"grant_type": []string{"client_credentials"},
		"scope":      []string{"https://api.ebay.com/oauth/api_scope"},
	}.Encode()

	// Формируем запрос на авторизацию
	r, err := http.NewRequestWithContext(ctx, http.MethodPost, authPath, bytes.NewBuffer([]byte(body)))
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	// Добавляем заголовок авторизации
	r.Header.Set("Authorization", fmt.Sprintf("Basic %s", req.BasicToken))

	// Добавляем заголовок Content-Type
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// Делаем запрос
	resp, err := s.httpClient.Do(r)
	if err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	// Проверяем код ответа
	if resp.StatusCode != http.StatusOK {
		return res, errors.InternalServer.New(fmt.Sprintf("Auth failed with http code %d", resp.StatusCode))
	}

	// Читаем ответ
	if err = json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return res, errors.InternalServer.Wrap(err)
	}

	return res, nil
}
