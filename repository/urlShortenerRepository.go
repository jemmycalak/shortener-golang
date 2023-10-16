package repository

import (
	"context"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"jemmy-sapta/model"
	"jemmy-sapta/protos/pbuild"
)

type ShortenerRepository interface {
	GetLastId() int64
	Update(*pbuild.UrlShortenerProto) (pbuild.UrlShortenerProto, error)
	SearchById(int64) (*model.Shortener, error)
	SearchByShortUrl(string) (*pbuild.UrlShortenerProto, error)
	Delete(int) error
	Create(*pbuild.UrlShortenerProto) (pbuild.UrlShortenerProto, error)

	IncreaseAccessed(shortener string, ipAddress string) error
	GetCountStatisticByDate(shortenerId int64, startDate string, endDate string) int64
	GetCountStatisticById(shortenerId int64) int64
	SearchStatisticBy(shortenerId int64, ipAddress string) (*model.StatisticShortener, error)

	SetShortenerCache(key string, value interface{}) error
	GetShortenerCache(key string) (string, error)
	IncreaseCounterCache(key string) error
	IsCacheEmpty() bool
}

type ResourceShortener struct {
	Db        *gorm.DB
	RDbClient *redis.Client
	Context   context.Context
}

func InitShortenerRepository(
	db *gorm.DB,
	rDbClient *redis.Client,
	context context.Context,
) *ResourceShortener {
	return &ResourceShortener{
		Db:        db,
		RDbClient: rDbClient,
		Context:   context,
	}
}

func (r *ResourceShortener) GetLastId() int64 {
	var data model.Shortener
	r.Db.Raw("SELECT id from shorteners order by id desc limit 1").Scan(&data)

	return data.Id
}

func (r *ResourceShortener) Create(data *pbuild.UrlShortenerProto) (pbuild.UrlShortenerProto, error) {
	lastInsertId := int64(0)
	errorInsert := r.Db.Raw("INSERT INTO shorteners (base_url, redirect, user_id) values ($1, $2, $3) RETURNING id",
		data.BaseUrl,
		data.Redirect,
		data.UserId,
	).Scan(&lastInsertId)

	if errorInsert.Error != nil {
		return pbuild.UrlShortenerProto{}, errorInsert.Error
	}

	searchData, errorSearch := r.SearchById(lastInsertId)
	if errorSearch != nil {
		return pbuild.UrlShortenerProto{}, errorSearch
	}

	data.Id = lastInsertId
	data.BaseUrl = searchData.BaseUrl

	return *data, errorInsert.Error
}

func (r *ResourceShortener) Update(data *pbuild.UrlShortenerProto) (pbuild.UrlShortenerProto, error) {
	var updatedData model.Shortener
	errorUpdate := r.Db.Raw("UPDATE shorteners SET base_url=?, redirect=? WHERE id=?",
		data.BaseUrl,
		data.Redirect,
		data.Id,
	).Scan(&updatedData)

	if errorUpdate.Error != nil {
		return *data, errorUpdate.Error
	}

	return *data, nil
}

func (r *ResourceShortener) SearchById(id int64) (*model.Shortener, error) {
	var data model.Shortener
	errorSearchQuery := r.Db.Raw("SELECT * FROM shorteners WHERE id=$1", id).Scan(&data)
	return &data, errorSearchQuery.Error
}

func (r *ResourceShortener) SearchByShortUrl(shortener string) (*pbuild.UrlShortenerProto, error) {
	var data model.Shortener
	errorSearchQuery := r.Db.Raw("SELECT * FROM shorteners WHERE base_url=$1", shortener).First(&data)

	if errorSearchQuery.Error != nil {
		return nil, errorSearchQuery.Error
	}

	return &pbuild.UrlShortenerProto{
		Id:       data.Id,
		BaseUrl:  data.BaseUrl,
		Redirect: data.Redirect,
		Clicked:  data.Clicked,
		UserId:   data.UserId,
	}, errorSearchQuery.Error
}

func (r *ResourceShortener) SearchStatisticBy(idShortener int64, ipAddress string) (*model.StatisticShortener, error) {
	var data model.StatisticShortener
	errorSearchQuery := r.Db.Raw(
		"SELECT * FROM statistic_shorteners WHERE shortener_id=$1 AND ip_address=$2",
		idShortener,
		ipAddress,
	).Scan(&data)

	if errorSearchQuery.Error != nil {
		return nil, errorSearchQuery.Error
	}

	return &data, errorSearchQuery.Error
}

func (r *ResourceShortener) Delete(id int) error {
	errorDeleteQuery := r.Db.Exec("DELETE FROM shorteners WHERE id=$1", id)

	return errorDeleteQuery.Error
}

func (r *ResourceShortener) IncreaseAccessed(shortener string, ipAddress string) error {
	data, errorSearch := r.SearchByShortUrl(shortener)
	if errorSearch != nil {
		return errorSearch
	}

	errorInsertQuery := r.Db.Exec("INSERT INTO statistic_shorteners (shortener_id, ip_address) VALUES ($1, $2)",
		data.Id,
		ipAddress,
	)

	if errorInsertQuery.Error != nil {
		return errorInsertQuery.Error
	}
	return nil
}

func (r *ResourceShortener) GetCountStatisticByDate(
	shortenerId int64,
	startDate string,
	endDate string,
) int64 {
	var count int64 = 0
	errorSearchStaticQuery := r.Db.Raw("SELECT COUNT(id) FROM statistic_shorteners WHERE shortener_id=$1 AND date BETWEEN $2 AND $3",
		shortenerId,
		startDate,
		endDate,
	).Scan(&count)

	if errorSearchStaticQuery.Error != nil {
		return count
	}
	return count
}

func (r *ResourceShortener) GetCountStatisticById(
	shortenerId int64,
) int64 {
	var count int64 = 0
	errorSearchStaticQuery := r.Db.Raw("SELECT COUNT(id) FROM statistic_shorteners WHERE shortener_id=$1",
		shortenerId,
	).Scan(&count)

	if errorSearchStaticQuery.Error != nil {
		return count
	}
	return count
}

func (r *ResourceShortener) SetShortenerCache(key string, value interface{}) error {
	errorSave := r.RDbClient.Set(r.Context, key, value, 0).Err()
	if errorSave != nil {
		return errorSave
	}
	return nil
}

func (r *ResourceShortener) GetShortenerCache(key string) (string, error) {
	result, errorQuery := r.RDbClient.Get(r.Context, key).Result()
	if errorQuery != nil {
		return "", errorQuery
	}
	return result, nil
}

func (r *ResourceShortener) IsCacheEmpty() bool {
	result, errorKeys := r.RDbClient.Keys(r.Context, "*").Result()
	if errorKeys != nil {
		return true
	}
	return len(result) == 0
}

func (r *ResourceShortener) IncreaseCounterCache(key string) error {
	return r.RDbClient.Incr(r.Context, key).Err()
}
