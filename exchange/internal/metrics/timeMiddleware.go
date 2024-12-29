package metrics

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2/utils"

	"github.com/gofiber/fiber/v2"

	"pkg/errors"
)

func ResponseTimeMiddleware(preparePathFunc func(string) string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		if globalMetrics.responseTimeMetric == nil {
			return errors.InternalServer.New("ResponseTime prometheus metric not initialized")
		}

		start := time.Now()

		path := utils.CopyString(c.Path())
		if preparePathFunc != nil {
			path = preparePathFunc(path)
		}

		// Сохраняем статус ответа после выполнения следующего в стеке хандлера
		err := c.Next()

		// Вычисляем продолжительность
		duration := time.Since(start)

		// Записываем информацию о времени ответа с использованием прометеуса
		globalMetrics.responseTimeMetric.WithLabelValues(
			preparePath(path),
			fmt.Sprintf("%d", c.Response().StatusCode()),
		).Observe(duration.Seconds())

		return err
	}
}

// preparePath заменяет пути /rtb/ssp на rtb_ssp
func preparePath(path string) string {
	path = strings.TrimPrefix(path, "/")
	return strings.ReplaceAll(path, "/", "_")
}
