package main

import (
	"context"
	"fmt"
	"jemmy-sapta/configs"
	"jemmy-sapta/server/urlshortenerconsumerkafka/service"
)

func main() {
	config, errorLoadConfig := configs.LoadConfig("./configs")
	if errorLoadConfig != nil {
		panic("error failed load app config")
	}
	rDbClient := configs.InitialRedisConnection(config, context.Background())
	grpcClient := configs.InitGrpcClient(config)

	clientService := grpcClient.GetShortenerClientService()
	shortenerService := service.Init(rDbClient, context.Background(), clientService)
	shortenerConsumerService := service.URlShortenerConsumerService(config, shortenerService, context.Background())
	shortenerConsumerService.FetchMessage()

	fmt.Println("Consumer is running")
}
