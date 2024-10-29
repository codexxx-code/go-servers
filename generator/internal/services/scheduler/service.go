package scheduler

import (
	"context"

	"github.com/robfig/cron/v3"

	generatorModel "generator/internal/services/generator/model"
	generatorService "generator/internal/services/generator/service"
)

type Scheduler struct {
	cron              *cron.Cron
	generatorService  GeneratorService
	generationEnabled bool
}

var _ GeneratorService = new(generatorService.GeneratorService)

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
