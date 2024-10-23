package scheduler

import (
	"context"

	"pkg/datetime"
	"pkg/log"
	"server/internal/services/forecast/model/zodiac"
	generatorModel "server/internal/services/generator/model"
)

func (s *Scheduler) makeDailyZodiacForecast(ctx context.Context) {

	todayDate := datetime.Now()

	// Проходимся по каждому знаку зодиака
	for _, zodiac := range zodiac.Array {

		// Генерируем прогноз для текущего знака зодиака на сегодня
		if err := s.generatorService.GenerateDailyZodiacForecast(ctx, generatorModel.GenerateDailyZodiacForecastReq{
			Date:   todayDate,
			Zodiac: zodiac,
		}); err != nil {
			log.Error(ctx, err)
		}
	}
}
