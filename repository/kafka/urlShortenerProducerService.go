package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"jemmy-sapta/configs"
	"jemmy-sapta/repository"
)

type ShortenerProducerService interface {
	TriggerUpdateAccessed(string) error
}

type ShortenerProducer struct {
	Writer              *kafka.Writer
	UrlShortenerService repository.ShortenerRepository
	Context             context.Context
}

func URlShortenerProducerService(
	config configs.AppConfig,
	urlShortenerService repository.ShortenerRepository,
	ctx context.Context,
) *ShortenerProducer {
	return &ShortenerProducer{
		Writer: &kafka.Writer{
			Addr:  kafka.TCP(config.APP_KAFKA_HOST),
			Topic: config.KAFKA_TOPIC,
		},
		UrlShortenerService: urlShortenerService,
		Context:             ctx,
	}
}

func (w *ShortenerProducer) TriggerUpdateAccessed(
	message string,
) error {
	data := []byte(message)
	errorWriteMessage := w.Writer.WriteMessages(w.Context, kafka.Message{
		Value: data,
	})
	if errorWriteMessage != nil {
		return errorWriteMessage
	}
	return nil
}
