package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/horoscope/repository/horoscopeDDL"
	"pkg/ddlHelper"

	"generator/internal/services/horoscope/model"
)

func (r *HoroscopeRepository) GetHoroscopes(ctx context.Context, req model.GetHoroscopesReq) (forecasts []model.Horoscope, err error) {

	filters := make(sq.Eq)

	if len(req.Dates) != 0 {
		filters[horoscopeDDL.ColumnDate] = req.Dates
	}

	if len(req.Zodiacs) != 0 {
		filters[horoscopeDDL.ColumnZodiac] = req.Zodiacs
	}

	// Выполняем запрос
	return forecasts, r.db.Select(ctx, &forecasts, sq.
		Select(ddlHelper.SelectAll).
		From(horoscopeDDL.Table).
		Where(filters),
	)
}
