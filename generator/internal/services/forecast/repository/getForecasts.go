package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"pkg/ddlHelper"

	"generator/internal/services/forecast/model"
)

func (r *ForecastRepository) GetForecasts(ctx context.Context, req model.GetForecastsReq) (forecasts []model.Forecast, err error) {

	filters := make(sq.Eq)

	if len(req.Dates) != 0 {
		filters["date"] = req.Dates
	}

	if len(req.Zodiacs) != 0 {
		filters["zodiac"] = req.Zodiacs
	}

	// Выполняем запрос
	return forecasts, r.db.Select(ctx, &forecasts, sq.
		Select(ddlHelper.SelectAll).
		From("zodiac.forecasts").
		Where(filters),
	)
}
