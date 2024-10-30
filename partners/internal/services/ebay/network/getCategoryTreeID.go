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
	getCategoryTreeIDPath = "/commerce/taxonomy/v1/get_default_category_tree_id"
)

func (n *EbayNetwork) GetCategoryTreeID(ctx context.Context, req model.GetCategoryTreeIDReq) (res model.GetCategoryTreeIDRes, err error) {

	// Подготавливаем query параметры
	urlValues := url.Values{
		"marketplace_id": []string{req.MarketplaceID},
	}

	// Отправляем запрос
	resp, err := n.sendRequest(sendRequestDTO{
		ctx:       ctx,
		method:    http.MethodGet,
		path:      getCategoryTreeIDPath,
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
