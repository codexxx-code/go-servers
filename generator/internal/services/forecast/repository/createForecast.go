package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/forecast/model"
)

func (r *ForecastRepository) CreateForecast(ctx context.Context, req model.CreateForecastReq) (id uint32, err error) {

	return r.db.ExecWithLastInsertID(ctx, sq.
		Insert("zodiac.forecasts").
		SetMap(map[string]any{
			"zodiac": req.Zodiac,
			"date":   req.Date,
			"text":   req.Text,
		}),
	)

}
