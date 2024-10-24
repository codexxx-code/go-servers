package scheduler

import (
	"context"

	"github.com/robfig/cron/v3"

	generatorModel "server/internal/services/generator/model"
)

type Scheduler struct {
	cron              *cron.Cron
	generatorService  GeneratorService
	generationEnabled bool
}

type GeneratorService interface {
	GenerateDailyZodiacForecast(ctx context.Context, req generatorModel.GenerateDailyZodiacForecastReq) error
}

func NewScheduler(
	generatorService GeneratorService,
	generationEnabled bool,
) *Scheduler {
	return &Scheduler{
		generatorService:  generatorService,
		cron:              cron.New(),
		generationEnabled: generationEnabled,
	}
}
