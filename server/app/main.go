package main

import (
	"context"
	"fmt"
	"jemmy-sapta/api"
	"jemmy-sapta/configs"
	"jemmy-sapta/middleware"
	"jemmy-sapta/repository"
	"jemmy-sapta/repository/kafka"
)

func main() {
	config, errorLoadConfig := configs.LoadConfig("./configs")
	if errorLoadConfig != nil {
		panic("error failed load app config")
	}

	fmt.Println("It's awasome !")
	db := configs.InitialDatabase(config)
	rDbClient := configs.InitialRedisConnection(config, context.Background())
	grpcClient := configs.InitGrpcClient(config)

	// initial repository
	shortenerService := repository.InitShortenerRepository(db, rDbClient, context.Background())
	producerService := kafka.URlShortenerProducerService(config, shortenerService, context.Background())
	shortenerClientService := grpcClient.GetShortenerClientService()
	userClientService := grpcClient.GetUserClientService()

	middleware.InitMiddlewareApi(config, userClientService)
	api.InitShortenerApi(producerService, shortenerClientService)
	api.InitUserApi(userClientService)
	api.InitialRoute(config)
}
