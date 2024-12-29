package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/horoscope/repository/horoscopeDDL"
	"pkg/ddlHelper"

	"generator/internal/services/horoscope/model"
)

func (r *HoroscopeRepository) GetHoroscope(ctx context.Context, req model.GetHoroscopeReq) (horoscope model.Horoscope, err error) {

	filters := sq.Eq{
		horoscopeDDL.ColumnTimeframe:     req.Timeframe,
		horoscopeDDL.ColumnPrimaryZodiac: req.PrimaryZodiac,
		horoscopeDDL.ColumnLanguage:      req.Language,
		horoscopeDDL.ColumnHoroscopeType: req.HoroscopeType,
	}

	if req.SecondaryZodiac != nil {
		filters[horoscopeDDL.ColumnSecondaryZodiac] = req.SecondaryZodiac
	}

	q := sq.Select(ddlHelper.SelectAll).
		From(horoscopeDDL.Table).
		Where(filters).
		Where(sq.LtOrEq{horoscopeDDL.ColumnDateFrom: req.DateFrom}).
		Where(sq.GtOrEq{horoscopeDDL.ColumnDateTo: req.DateFrom}).
		OrderBy(ddlHelper.Desc(horoscopeDDL.ColumnID)).
		Limit(1)

	// Выполняем запрос
	return horoscope, r.db.Get(ctx, &horoscope, q)
}
