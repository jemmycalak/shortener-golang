package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"jemmy-sapta/configs"
	"jemmy-sapta/middleware"
	"jemmy-sapta/protos/pbuild"
	"jemmy-sapta/repository"
	"jemmy-sapta/usecase"
	"net"
)

func main() {
	config, errorLoadConfig := configs.LoadConfig("./configs")
	if errorLoadConfig != nil {
		panic("error failed load app config")
	}

	listener, errorLoadConfig := net.Listen("tcp", config.APP_GRPC_HOST)
	if errorLoadConfig != nil {
		fmt.Printf("Failed to listen: %v\n", errorLoadConfig)
	}

	db := configs.InitialDatabase(config)
	rDbClient := configs.InitialRedisConnection(config, context.Background())

	shortenerRepository := repository.InitShortenerRepository(db, rDbClient, context.Background())
	userRepository := repository.InitUserRepository(db)
	usecase.InitShortenerUsercase(shortenerRepository)
	usecase.InitUserUsercase(userRepository)
	userUsecase := usecase.UserUsecase{}

	grpcInterceptor := middleware.GrpcInterceptor{
		UserUsecase:   userUsecase,
		AppConfig:     config,
		AllowedMethod: allowedRoles(),
	}
	server := grpc.NewServer(
		grpc.UnaryInterceptor(grpcInterceptor.GetUnaryInterceptor()),
	)
	pbuild.RegisterUrlShortenerServiceProtoServer(
		server,
		&usecase.UrlShortenerUsecase{},
	)
	pbuild.RegisterUserServiceProtoServer(
		server,
		&userUsecase,
	)

	reflection.Register(server)

	if errorServe := server.Serve(listener); errorServe != nil {
		fmt.Printf("Failed to serve: %v", errorServe)
	}
}

func allowedRoles() map[string][]string {
	const userServicePath = "/protos.UserServiceProto/"
	const shortenerServicePath = "/protos.UrlShortenerServiceProto/"
	return map[string][]string{
		userServicePath + "GetUserBy":             {"user"},
		userServicePath + "CreateUser":            {"user"},
		shortenerServicePath + "SearchByShortUrl": {"user"},
		shortenerServicePath + "IncreaseAccessed": {"user"},
	}
}
