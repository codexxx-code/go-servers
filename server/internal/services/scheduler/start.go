package scheduler

import (
	"context"
	"time"

	"pkg/errors"
)

func (s *Scheduler) Start() error {

	// Создание ежедневного прогноза для знаков зодиака
	_, err := s.cron.AddFunc("@daily", func() { // Every day at 00:00 UTC

		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		defer cancel()

		s.makeDailyZodiacForecast(ctx)
	})
	if err != nil {
		return errors.InternalServer.Wrap(err)
	}

	s.cron.Start()

	return nil
}
