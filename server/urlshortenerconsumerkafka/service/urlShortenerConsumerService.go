package service

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
	"jemmy-sapta/configs"
)

type ShortenerConsumer struct {
	Reader              *kafka.Reader
	UrlShortenerService UrlShortenerService
	Context             context.Context
}

func URlShortenerConsumerService(
	config configs.AppConfig,
	urlShortenerService UrlShortenerService,
	ctx context.Context,
) *ShortenerConsumer {
	return &ShortenerConsumer{
		Reader: kafka.NewReader(kafka.ReaderConfig{
			Brokers:     []string{config.APP_KAFKA_HOST},
			Topic:       config.KAFKA_TOPIC,
			GroupID:     config.KAFKA_TOPIC,
			StartOffset: kafka.LastOffset,
		}),
		Context:             ctx,
		UrlShortenerService: urlShortenerService,
	}
}

func (r *ShortenerConsumer) FetchMessage() {
	for {
		message, errorFetchMessage := r.Reader.FetchMessage(r.Context)
		if errorFetchMessage != nil {
			fmt.Println("Error fetch message:", errorFetchMessage)
		}

		if message.Value != nil {
			r.UrlShortenerService.IncreaseAccessed(string(message.Value))
		}
	}
}
