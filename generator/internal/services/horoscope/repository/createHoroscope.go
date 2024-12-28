package repository

import (
	"context"

	sq "github.com/Masterminds/squirrel"

	"generator/internal/services/horoscope/model"
	"generator/internal/services/horoscope/repository/horoscopeDDL"
)

func (r *HoroscopeRepository) CreateHoroscope(ctx context.Context, req model.CreateHoroscopeReq) (id uint32, err error) {

	return r.db.ExecWithLastInsertID(ctx, sq.
		Insert(horoscopeDDL.Table).
		SetMap(map[string]any{
			horoscopeDDL.ColumnDateFrom:        req.DateFrom,
			horoscopeDDL.ColumnDateTo:          req.DateTo,
			horoscopeDDL.ColumnPrimaryZodiac:   req.PrimaryZodiac,
			horoscopeDDL.ColumnSecondaryZodiac: req.SecondaryZodiac,
			horoscopeDDL.ColumnLanguage:        req.Language,
			horoscopeDDL.ColumnTimeframe:       req.Timeframe,
			horoscopeDDL.ColumnHoroscopeType:   req.HoroscopeType,
			horoscopeDDL.ColumnText:            req.Text,
		}),
	)
}
