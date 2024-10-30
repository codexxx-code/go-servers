package network

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"partners/internal/services/ebay/network/model"
	"pkg/errors"
)

const (
	itemsPath = "/buy/browse/v1/item_summary/search"
)

func (n *EbayNetwork) GetItems(ctx context.Context, req model.GetItemsReq) (res model.GetItemsRes, err error) {

	// Формируем параметры запроса
	var urlValues = make(url.Values)

	// Добавляем фильтр по идентификаторам категорий
	if req.CategoryID != nil {
		urlValues.Add("category_ids", *req.CategoryID)
	}

	// Отправляем запрос
	resp, err := n.sendRequest(sendRequestDTO{
		ctx:       ctx,
		method:    http.MethodGet,
		path:      itemsPath,
		urlValues: urlValues,
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
