package network

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"partners/internal/services/ebay/network/model"
	"pkg/errors"
)

const (
	categoriesPath = "/commerce/taxonomy/v1/category_tree/%s"
)

func (n *EbayNetwork) GetCategories(ctx context.Context, req model.GetCategoriesReq) (res model.GetCategoriesRes, err error) {

	// Подготавливаем путь
	path := fmt.Sprintf(categoriesPath, req.CategoryTreeID)

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
