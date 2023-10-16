package configs

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"jemmy-sapta/client/service"
	"jemmy-sapta/protos/pbuild"
)

type ResourceClient struct {
	Client *grpc.ClientConn
}

type GrpcClient interface {
	GetShortenerClientService() service.ShortenerClientService
	GetUserClientService() service.UserClientService
}

func InitGrpcClient(config AppConfig) *ResourceClient {
	// connect to server gRPC
	client, errorConnectGrPC := grpc.Dial(
		config.APP_GRPC_HOST,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if errorConnectGrPC != nil {
		fmt.Printf("gagal menyambungkan ke server: %v", errorConnectGrPC)
		return nil
	}

	return &ResourceClient{
		Client: client,
	}
}

func (r *ResourceClient) GetShortenerClientService() service.ShortenerClientService {
	var clientService = pbuild.NewUrlShortenerServiceProtoClient(r.Client)
	return service.ShortenerClientService{Client: clientService}
}

func (r *ResourceClient) GetUserClientService() service.UserClientService {
	var clientService = pbuild.NewUserServiceProtoClient(r.Client)
	return service.UserClientService{Client: clientService}
}
