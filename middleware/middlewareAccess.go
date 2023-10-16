package middleware

import (
	"context"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc/metadata"
	"jemmy-sapta/client/service"
	"jemmy-sapta/configs"
	"jemmy-sapta/protos/pbuild"
)

var appConfig configs.AppConfig

var userClientService service.UserClientService

func InitMiddlewareApi(
	config configs.AppConfig,
	clientService service.UserClientService,
) {
	userClientService = clientService
	appConfig = config
}

func CreateHandlerConfig() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(appConfig.SECRET_KEY)},
	})
}

func IsValidAuthToken(c *fiber.Ctx) error {
	accessToken := GetToken(c)
	newContext := CreateGrpcContext(c.Context(), accessToken)
	if _, errorValidateToken := userClientService.Client.AutorizeToken(newContext, &pbuild.AuthorizeTokenRequest{
		AccessToken: accessToken,
	}); errorValidateToken != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "fail", "message": errorValidateToken.Error()})
	}

	return c.Next()
}

func CreateGrpcContext(
	ctx context.Context,
	accessToken string,
) context.Context {
	// create header and new context to hit grpc
	headerGrpc := metadata.New(map[string]string{"authorization": accessToken})
	return metadata.NewOutgoingContext(ctx, headerGrpc)
}
