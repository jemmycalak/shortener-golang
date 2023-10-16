package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"jemmy-sapta/client/service"
	"jemmy-sapta/protos/pbuild"
	"log"
)

func main() {
	// menyambungkan ke server gRPC
	client, errorConnectGrPc := grpc.Dial("localhost:1446", grpc.WithInsecure())
	if errorConnectGrPc != nil {
		log.Fatalf("gagal menyambungkan ke server grpc: %v", errorConnectGrPc)
	}

	var clientService = pbuild.NewUrlShortenerServiceProtoClient(client)

	services := service.ShortenerClientService{Client: clientService}

	// call proto service client
	response, errorResponse := services.Client.Create(context.Background(), &pbuild.UrlShortenerProto{
		BaseUrl: "Example Request",
	})

	if errorResponse != nil {
		fmt.Printf("Error Respnse %v", errorResponse)
	}
	fmt.Println("Response ::", response)
}
