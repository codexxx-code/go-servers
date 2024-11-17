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
	itemsSummaryPath = "/buy/browse/v1/item_summary/search"
)

func (n *EbayNetwork) GetItemsSummary(ctx context.Context, req model.GetItemsSummaryReq) (res model.GetItemsSummaryRes, err error) {

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
		path:      itemsSummaryPath,
		urlValues: urlValues,
		headers: map[string]string{
			epnHeader: fmt.Sprintf("affiliateCampaignId=%s", n.campaignID),
		},
		withAuth: true,
		body:     nil,
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
