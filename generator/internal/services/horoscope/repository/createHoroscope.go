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
			horoscopeDDL.ColumnZodiac: req.Zodiac,
			horoscopeDDL.ColumnDate:   req.Date,
			horoscopeDDL.ColumnText:   req.Text,
		}),
	)

}
