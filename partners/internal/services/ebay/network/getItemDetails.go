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

const (
	itemDetailsPath = "/buy/browse/v1/item/%s"
)

func (n *EbayNetwork) GetItemDetails(ctx context.Context, req model.GetItemDetailsReq) (res model.GetItemDetailsRes, err error) {

	// Подготавливаем путь
	req.ID = url.PathEscape(req.ID)
	path := fmt.Sprintf(itemDetailsPath, req.ID)

	// Отправляем запрос
	resp, err := n.sendRequest(sendRequestDTO{
		ctx:       ctx,
		method:    http.MethodGet,
		path:      path,
		urlValues: nil,
		headers:   nil,
		withAuth:  true,
		body:      nil,
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
