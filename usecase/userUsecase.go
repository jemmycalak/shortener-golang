package usecase

import (
	"context"
	"jemmy-sapta/protos/pbuild"
	"jemmy-sapta/repository"
)

var userRepository repository.UserRepository

type UserUsecase struct {
	pbuild.UnimplementedUserServiceProtoServer
}

func InitUserUsercase(repository repository.UserRepository) {
	userRepository = repository
}

func (*UserUsecase) CreateUser(
	ctx context.Context,
	request *pbuild.CreateUserRequest,
) (*pbuild.UserProto, error) {
	var body = pbuild.UserProto{
		Username: request.Username,
		Password: request.Password,
	}
	return userRepository.CreateUser(body)
}

func (*UserUsecase) GetUserBy(
	ctx context.Context,
	request *pbuild.GetUserByRequest,
) (*pbuild.UserProto, error) {
	return userRepository.GetUserBy(request.Username, request.Password)
}

func (*UserUsecase) GetUserById(
	ctx context.Context,
	request *pbuild.GetUserByIdRequest,
) (*pbuild.UserProto, error) {
	return userRepository.GetUserById(request.Id)
}

func (*UserUsecase) AutorizeToken(
	context.Context,
	*pbuild.AuthorizeTokenRequest,
) (*pbuild.AuthorizeTokenResponse, error) {
	return &pbuild.AuthorizeTokenResponse{}, nil
}
