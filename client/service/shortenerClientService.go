package service

import (
	"jemmy-sapta/protos/pbuild"
)

type ShortenerClientService struct {
	Client pbuild.UrlShortenerServiceProtoClient
}
