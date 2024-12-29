package repository

import (
	"context"

	"github.com/Shopify/sarama"

	saramaHelper "pkg/kafka"
)

type AnalyticWriterRepository struct {
	kafka sarama.AsyncProducer
}

func NewAnalyticWriterRepository(
	kafka sarama.AsyncProducer,
) *AnalyticWriterRepository {
	return &AnalyticWriterRepository{
		kafka: kafka,
	}
}

// writeToTopic записывает сообщение в топик и обрабатывает возможные ошибки
func (r *AnalyticWriterRepository) writeToTopic(_ context.Context, message any, topic string) error {

	m, err := saramaHelper.ConvertStructToMessage(message, topic)
	if err != nil {
		return err
	}

	// Отправляем сообщения
	r.kafka.Input() <- m

	return nil
}
