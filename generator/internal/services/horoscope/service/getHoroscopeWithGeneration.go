package service

import (
	"context"
	"fmt"

	generatorModel "generator/internal/services/generator/model"
	"generator/internal/services/horoscope/model"
	"generator/internal/services/horoscope/model/zodiac"
	"pkg/datetime"
	"pkg/errors"
	"pkg/log"
	"pkg/slices"
)

func (s *HoroscopeService) GetHoroscopeWithGeneration(ctx context.Context, req model.GetHoroscopeWithGenerationReq) (res model.Horoscope, err error) {

	// Получаем гороскоп
	horoscope, err := slices.FirstWithError(
		s.horoscopeRepository.GetHoroscopes(ctx, model.GetHoroscopesReq{
			Dates:   []datetime.Date{req.Date},
			Zodiacs: []zodiac.Zodiac{req.Zodiac},
		}),
	)
	if err != nil {

		// Если гороскоп не найден
		if errors.Is(err, slices.ErrSliceIsEmpty) {

			log.Info(ctx, fmt.Sprintf("Не нашли гороскоп для знака зодиака %s в дату %s, генерируем новый", req.Zodiac.ToRussian(), req.Date.String()))

			// Генерируем новый гороскоп
			generationRes, err := s.generatorService.GenerateDailyHoroscope(ctx, generatorModel.GenerateDailyHoroscopeReq{
				Date:   req.Date,
				Zodiac: req.Zodiac,
			})
			if err != nil {
				return res, err
			}

			// Добавляем гороскоп в базу данных
			id, err := s.horoscopeRepository.CreateHoroscope(ctx, model.CreateHoroscopeReq{
				Date:   req.Date,
				Zodiac: req.Zodiac,
				Text:   generationRes.Text,
			})
			if err != nil {
				return res, err
			}

			// Возвращаем гороскоп
			return model.Horoscope{
				ID:     id,
				Date:   req.Date,
				Zodiac: req.Zodiac,
				Text:   generationRes.Text,
			}, nil

		}

		return res, err
	}

	return horoscope, nil

}
