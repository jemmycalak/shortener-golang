package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"jemmy-sapta/configs"
	"jemmy-sapta/protos/pbuild"
	"jemmy-sapta/usecase"
)

type GrpcInterceptor struct {
	UserUsecase   usecase.UserUsecase
	AppConfig     configs.AppConfig
	AllowedMethod map[string][]string
}

func (gInterceptor *GrpcInterceptor) GetUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		errorAuthorization := gInterceptor.authorizeUnaryInterceptor(ctx, info)
		if errorAuthorization != nil {
			return nil, errorAuthorization
		}
		return handler(ctx, req)
	}
}

func (gInterceptor *GrpcInterceptor) authorizeUnaryInterceptor(
	ctx context.Context,
	info *grpc.UnaryServerInfo,
) error {
	_, isAllowedMethod := gInterceptor.AllowedMethod[info.FullMethod]
	if isAllowedMethod {
		return nil
	}

	metaData, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "kamu harus login dahulu")
	}

	authorization := metaData["authorization"]
	if len(authorization) == 0 {
		return status.Errorf(codes.Unauthenticated, "dibutuhkan token authorization")
	}

	accessToken := authorization[0]

	if _, errorValidateToken := gInterceptor.IsValidToken(ctx, accessToken); errorValidateToken != nil {
		return status.Errorf(codes.Unauthenticated, errorValidateToken.Error())
	}

	return nil
}

func (gInterceptor *GrpcInterceptor) IsValidToken(
	ctx context.Context,
	accessToken string,
) (interface{}, error) {
	if accessToken == "" {
		return nil, fmt.Errorf("you are not logged in")
	}

	tokenByte, err := jwt.Parse(accessToken, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(gInterceptor.AppConfig.SECRET_KEY), nil
	})
	if err != nil {
		return nil, fmt.Errorf("invalidate token: %v", err)
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return nil, fmt.Errorf("invalid token claim")
	}

	userData, errorGetData := gInterceptor.UserUsecase.GetUserById(ctx, &pbuild.GetUserByIdRequest{
		Id: int64(claims["userId"].(float64)),
	})

	if errorGetData != nil {
		return nil, fmt.Errorf("user not found")
	}
	if !IsAdmin(claims) {
		return nil, fmt.Errorf("your are not admin user")
	}

	return userData.Id, nil
}

func IsAdmin(claims jwt.MapClaims) bool {
	return claims["admin"].(bool)
}
