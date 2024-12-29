package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"exchange/internal/services/currency/model"
	"exchange/internal/services/currency/repository/currencyDDL"
	"pkg/ddlHelper"
)

func (r *CurrencyRepository) GetCurrencies(ctx context.Context) (currencies []model.Currency, err error) {

	// Получаем данные из кэша
	currencies, ok := r.cache.Get()
	if !ok {

		// Получаем данные из БД
		if err = r.pgsql.Select(ctx, &currencies, sq.
			Select(ddlHelper.SelectAll).
			From(currencyDDL.Table),
		); err != nil {
			return nil, err
		}
	}

	return currencies, nil
}
