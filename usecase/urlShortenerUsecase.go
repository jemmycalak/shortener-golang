package usecase

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/sourcegraph/conc"
	"golang.org/x/sync/errgroup"
	"google.golang.org/protobuf/types/known/emptypb"
	"jemmy-sapta/protos/pbuild"
	"jemmy-sapta/repository"
	"strconv"
	"strings"
	"time"
)

var shortenerRepository repository.ShortenerRepository

type UrlShortenerUsecase struct {
	pbuild.UnimplementedUrlShortenerServiceProtoServer
}

const (
	ShortenerKey = "ShortenerKey:"
	DAILY        = ":1D"
	WEEKLY       = ":1W"
	MONTHLY      = ":1M"
	ALLTIME      = ":AT"
)

func InitShortenerUsercase(respository repository.ShortenerRepository) {
	shortenerRepository = respository
}

func (*UrlShortenerUsecase) Create(
	ctx context.Context,
	body *pbuild.UrlShortenerProto,
) (*pbuild.UrlShortenerProto, error) {
	var eg errgroup.Group
	var result *pbuild.UrlShortenerProto

	eg.Go(func() error {
		var id = shortenerRepository.GetLastId()
		var lengthId int
		if len(strconv.FormatInt(id, 10)) > 5 {
			lengthId = 5
		} else {
			lengthId = len(strconv.FormatInt(id, 10))
		}

		body.BaseUrl = RandomURL(lengthId)
		body.Redirect = IsContainHttps(body.Redirect)
		mResult, errorCreate := shortenerRepository.Create(body)
		if errorCreate != nil {
			return errorCreate
		}
		result = &mResult

		return nil
	})

	if errorEdit := eg.Wait(); errorEdit != nil {
		return nil, errors.New(errorEdit.Error())
	}
	return result, nil
}

func (*UrlShortenerUsecase) Edit(
	ctx context.Context,
	body *pbuild.UrlShortenerProto,
) (*pbuild.UrlShortenerProto, error) {
	var eg errgroup.Group
	var data *pbuild.UrlShortenerProto

	eg.Go(func() error {
		shortenerData, errorSearch := shortenerRepository.SearchById(body.Id)
		if errorSearch != nil {
			return errors.New("shortener data not found")
		}
		// cek if url is main
		if shortenerData != nil && shortenerData.UserId != body.UserId {
			return errors.New("shortener data not yours")
		}

		// cek if url is already exist
		_, errorIsExist := shortenerRepository.SearchByShortUrl(body.BaseUrl)
		if errorIsExist == nil {
			return errors.New("can not duplicate shortener")
		}

		body.Redirect = IsContainHttps(body.Redirect)
		dataUpdated, errorUpdate := shortenerRepository.Update(body)
		if errorUpdate != nil {
			return errorUpdate
		}
		data = &dataUpdated

		return nil
	})
	if errorEdit := eg.Wait(); errorEdit != nil {
		return nil, errors.New(errorEdit.Error())
	}

	return data, nil
}

func (usecase *UrlShortenerUsecase) Statistic(
	ctx context.Context,
	body *pbuild.GetStatisticByUrlRequest,
) (*pbuild.GetStatisticByUrlResponse, error) {
	data, errorSearch := shortenerRepository.SearchByShortUrl(body.BaseUrl)
	if errorSearch != nil {
		return nil, errorSearch
	}
	if data.UserId != body.UserId {
		return nil, errors.New("shortener data not yours")
	}

	dailyKey := ShortenerKey + data.BaseUrl + DAILY
	weeklyKey := ShortenerKey + data.BaseUrl + WEEKLY
	monthlyKey := ShortenerKey + data.BaseUrl + MONTHLY
	allTimeKey := ShortenerKey + data.BaseUrl + ALLTIME

	dateNow := time.Now()
	var daily int64
	var weekly int64
	var monthly int64
	var allTime int64

	var eg errgroup.Group
	eg.Go(func() error {
		daily = usecase.getCountStatistic(
			dailyKey,
			data.Id,
			dateNow.Format("2006-01-02"),
			dateNow.Format("2006-01-02"),
		)

		return nil
	})
	eg.Go(func() error {
		weekly = usecase.getCountStatistic(
			weeklyKey,
			data.Id,
			dateNow.AddDate(0, 0, -7).Format("2006-01-02"),
			dateNow.Format("2006-01-02"),
		)
		return nil
	})
	eg.Go(func() error {
		monthly = usecase.getCountStatistic(
			monthlyKey,
			data.Id,
			dateNow.AddDate(0, -1, 0).Format("2006-01-02"),
			dateNow.Format("2006-01-02"),
		)
		return nil
	})

	eg.Go(func() error {
		allTime = usecase.getCountStatistic(
			allTimeKey,
			data.Id,
			"",
			"",
		)
		return nil
	})

	if err := eg.Wait(); err != nil {
		return nil, err
	}

	return &pbuild.GetStatisticByUrlResponse{
		AllTime: allTime,
		Day:     daily,
		Week:    weekly,
		Month:   monthly,
	}, nil
}

func (*UrlShortenerUsecase) getCountStatistic(
	key string,
	shortenerId int64,
	startDate string,
	endDate string,
) int64 {
	var total int64 = 0
	cacheResult, errorGetShortenerCache := shortenerRepository.GetShortenerCache(key)

	if errorGetShortenerCache != nil && !strings.Contains(key, ALLTIME) {
		total = shortenerRepository.GetCountStatisticByDate(
			shortenerId,
			startDate,
			endDate,
		)

		shortenerRepository.SetShortenerCache(key, total)
	}
	if errorGetShortenerCache != nil && strings.Contains(key, ALLTIME) {
		total = shortenerRepository.GetCountStatisticById(shortenerId)
		shortenerRepository.SetShortenerCache(key, total)
	}

	if errorGetShortenerCache == nil {
		total = unMarshal(cacheResult, errorGetShortenerCache)
	}

	return total
}

func unMarshal(value string, errorCache error) int64 {
	var dataCahe int64 = 0
	if errorCache != nil {
		return dataCahe
	}

	if errorMarshal := json.Unmarshal([]byte(value), &dataCahe); errorMarshal != nil {
		return dataCahe
	}
	return dataCahe
}

func (*UrlShortenerUsecase) SearchByShortUrl(
	ctx context.Context,
	body *pbuild.SearchByShortUrlRequest,
) (*pbuild.UrlShortenerProto, error) {
	return shortenerRepository.SearchByShortUrl(body.BaseUrl)
}

func (usecase *UrlShortenerUsecase) IncreaseAccessed(
	ctx context.Context,
	body *pbuild.IncreaseAccessedRequest,
) (*emptypb.Empty, error) {
	shortener, errorSearch := shortenerRepository.SearchByShortUrl(body.BaseUrl)
	if errorSearch != nil {
		return &emptypb.Empty{}, errorSearch
	}
	statistic, errorSearchStatistic := shortenerRepository.SearchStatisticBy(shortener.Id, body.IpAddress)

	fmt.Println(shortener, "-------", statistic)
	if errorSearchStatistic != nil {
		var wg conc.WaitGroup
		var errorIncrease error

		wg.Go(func() {
			if errorIncreaseAccessed := shortenerRepository.IncreaseAccessed(body.BaseUrl, body.IpAddress); errorIncreaseAccessed != nil {
				errorIncrease = errorIncreaseAccessed
			}
		})

		wg.Go(func() {
			dailyKey := ShortenerKey + body.BaseUrl + DAILY
			weeklyKey := ShortenerKey + body.BaseUrl + WEEKLY
			monthlyKey := ShortenerKey + body.BaseUrl + MONTHLY
			allTimeKey := ShortenerKey + body.BaseUrl + ALLTIME
			usecase.increaseCounterCache(dailyKey)
			usecase.increaseCounterCache(weeklyKey)
			usecase.increaseCounterCache(monthlyKey)
			usecase.increaseCounterCache(allTimeKey)
		})
		wg.Wait()

		if errorIncrease != nil {
			return &emptypb.Empty{}, errorIncrease
		}
	}

	return &emptypb.Empty{}, nil
}

func (*UrlShortenerUsecase) increaseCounterCache(key string) error {
	var wg conc.WaitGroup
	var errorIncrease error

	wg.Go(func() {
		if errorIncrease = shortenerRepository.IncreaseCounterCache(key); errorIncrease != nil {
			fmt.Println(errorIncrease)
		}
	})
	wg.Wait()

	return errorIncrease
}
