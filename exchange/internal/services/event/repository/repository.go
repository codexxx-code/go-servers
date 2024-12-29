package repository

import (
	"context"

	"github.com/Shopify/sarama"
	"go.mongodb.org/mongo-driver/mongo"

	saramaHelper "pkg/kafka"
)

type EventRepository struct {
	kafka sarama.AsyncProducer
	mongo *mongo.Database
}

func NewEventRepository(
	kafka sarama.AsyncProducer,
	mongo *mongo.Database,
) *EventRepository {
	return &EventRepository{
		kafka: kafka,
		mongo: mongo,
	}
}

// writeToTopic записывает сообщение в топик и обрабатывает возможные ошибки
func (r *EventRepository) writeToTopic(_ context.Context, message any, topic string) error {

	m, err := saramaHelper.ConvertStructToMessage(message, topic)
	if err != nil {
		return err
	}

	r.kafka.Input() <- m

	return nil
}
