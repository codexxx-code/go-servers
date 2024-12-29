package repository

import (
	"github.com/Shopify/sarama"

	saramaHelper "pkg/kafka"
)

type BillingRepository struct {
	producer sarama.AsyncProducer
}

func NewBillingRepository(producer sarama.AsyncProducer) *BillingRepository {
	return &BillingRepository{
		producer: producer,
	}
}

// writeToTopic записывает сообщение в топик
func (r *BillingRepository) writeToTopic(message any, topic string) error {

	m, err := saramaHelper.ConvertStructToMessage(message, topic)
	if err != nil {
		return err
	}

	r.producer.Input() <- m
	return nil
}
