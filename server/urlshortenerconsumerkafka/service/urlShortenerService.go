package service

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"jemmy-sapta/client/service"
	"jemmy-sapta/model"
	"jemmy-sapta/protos/pbuild"
)

type UrlShortenerService interface {
	SearchByShortUrl(string) (*model.Shortener, error)
	IncreaseAccessed(message string) error
}

type Resource struct {
	Db                     *gorm.DB
	RdbClient              *redis.Client
	Context                context.Context
	ShortenerClientService service.ShortenerClientService
}

func Init(
	rDbClient *redis.Client,
	context context.Context,
	clientService service.ShortenerClientService,
) *Resource {
	return &Resource{
		RdbClient:              rDbClient,
		Context:                context,
		ShortenerClientService: clientService,
	}
}

func (r *Resource) SearchByShortUrl(shortener string) (*model.Shortener, error) {
	dataSearch, errorSearch := r.ShortenerClientService.Client.SearchByShortUrl(
		r.Context,
		&pbuild.SearchByShortUrlRequest{
			BaseUrl: shortener,
		},
	)
	data := model.Shortener{
		Id:       dataSearch.Id,
		BaseUrl:  dataSearch.BaseUrl,
		Redirect: dataSearch.Redirect,
		Clicked:  dataSearch.Clicked,
		UserId:   dataSearch.UserId,
	}

	if errorSearch != nil {
		return nil, errorSearch
	}

	return &data, errorSearch
}

func (r *Resource) IncreaseAccessed(message string) error {
	var triggerData model.TriggerShortener
	if errorUnMarshal := json.Unmarshal([]byte(message), &triggerData); errorUnMarshal != nil {
		return errorUnMarshal
	}

	_, errorUpdateQuery := r.ShortenerClientService.Client.IncreaseAccessed(
		r.Context,
		&pbuild.IncreaseAccessedRequest{
			BaseUrl:   triggerData.Url,
			IpAddress: triggerData.IpAddress,
		})

	if errorUpdateQuery != nil {
		return errorUpdateQuery
	}
	return nil
}
