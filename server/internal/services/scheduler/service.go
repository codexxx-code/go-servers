package scheduler

import (
	"context"

	"github.com/robfig/cron/v3"

	generatorModel "server/internal/services/generator/model"
)

type Scheduler struct {
	cron             *cron.Cron
	generatorService GeneratorService
}

type GeneratorService interface {
	GenerateDailyZodiacForecast(ctx context.Context, req generatorModel.GenerateDailyZodiacForecastReq) error
}

func NewScheduler(
	generatorService GeneratorService,
) *Scheduler {
	return &Scheduler{
		generatorService: generatorService,
		cron:             cron.New(),
	}
}
