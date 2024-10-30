package network

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"pkg/errors"
	"pkg/log"
)

type sendRequestDTO struct {
	ctx       context.Context
	method    string
	path      string
	headers   map[string]string
	urlValues url.Values
	withAuth  bool
	body      []byte
}

func (n *EbayNetwork) sendRequest(dto sendRequestDTO) (*http.Response, error) {

	// Формируем запрос на авторизацию
	r, err := http.NewRequestWithContext(dto.ctx, dto.method, n.baseHost+dto.path, bytes.NewBuffer(dto.body))
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Добавляем параметры запроса
	if len(dto.urlValues) != 0 {
		q := r.URL.Query()
		for key, value := range dto.urlValues {
			for _, v := range value {
				q.Add(key, v)
			}
		}
		r.URL.RawQuery = q.Encode()
	}

	// Добавляем заголовки
	for key, value := range dto.headers {
		r.Header.Set(key, value)
	}

	// При необходимости добавляем токен авторизации
	if dto.withAuth {

		// Получаем токен
		token, err := n.authManager.GetToken(dto.ctx)
		if err != nil {
			return nil, err
		}

		// Добавляем заголовок авторизации
		r.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	// Делаем запрос
	log.Info(dto.ctx, "Sending request", log.ParamsOption("url", r.URL.String()))
	resp, err := n.httpClient.Do(r)
	if err != nil {
		return nil, errors.InternalServer.Wrap(err)
	}

	// Проверяем код ответа
	switch {
	case resp.StatusCode >= 200 && resp.StatusCode <= 299:
		return resp, nil
	default:
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Error(dto.ctx, err)
		}
		return nil, errors.InternalServer.New(fmt.Sprintf("Request failed with http code %d", resp.StatusCode), errors.ParamsOption("body", string(body)))
	}
}
