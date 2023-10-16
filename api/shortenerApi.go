package api

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"jemmy-sapta/client/service"
	"jemmy-sapta/middleware"
	"jemmy-sapta/model"
	"jemmy-sapta/protos/pbuild"
	"jemmy-sapta/repository/kafka"
)

var shortenerProducerService kafka.ShortenerProducerService
var shortenerClientService service.ShortenerClientService

func InitShortenerApi(
	producerService kafka.ShortenerProducerService,
	clientService service.ShortenerClientService,
) {
	shortenerProducerService = producerService
	shortenerClientService = clientService
}

func wellcome(c *fiber.Ctx) error {
	return c.SendString("Wellcome")
}

func redirectUrl(c *fiber.Ctx) error {
	var url = c.Params("url")
	data, errorSearch := shortenerClientService.Client.SearchByShortUrl(
		c.Context(),
		&pbuild.SearchByShortUrlRequest{
			BaseUrl: url,
		},
	)

	if errorSearch != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "error " + errorSearch.Error(),
		})
	}
	triggerData := model.TriggerShortener{
		Url:       data.BaseUrl,
		IpAddress: c.IP(),
	}
	mData, errorMarshal := json.Marshal(triggerData)
	if errorMarshal != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "error " + errorMarshal.Error(),
		})
	}

	errorTriggered := shortenerProducerService.TriggerUpdateAccessed(string(mData))
	if errorTriggered != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error " + errorTriggered.Error(),
		})
	}

	return c.Redirect(data.Redirect, fiber.StatusMovedPermanently)
}

func getUserId(c *fiber.Ctx) int64 {
	tokenClaims := middleware.GetTokenClaims(c)
	userId := tokenClaims["userId"].(float64)

	return int64(userId)
}

func createUrlShortener(c *fiber.Ctx) error {
	c.Accepts("application/json")
	var body pbuild.UrlShortenerProto
	errorParsing := c.BodyParser(&body)
	if errorParsing != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing " + errorParsing.Error(),
		})
	}
	body.UserId = getUserId(c)
	newContext := middleware.CreateGrpcContext(c.Context(), middleware.GetToken(c))
	data, errorCreate := shortenerClientService.Client.Create(newContext, &body)
	if errorCreate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "create data error " + errorCreate.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "data created",
		"data":    data,
	})
}

func updateUrlShortener(c *fiber.Ctx) error {
	c.Accepts("application/json")

	var body pbuild.UrlShortenerProto
	errorParser := c.BodyParser(&body)
	if errorParser != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "error parsing " + errorParser.Error(),
		})
	}

	body.UserId = getUserId(c)
	newContext := middleware.CreateGrpcContext(c.Context(), middleware.GetToken(c))
	data, errorEdit := shortenerClientService.Client.Edit(newContext, &body)
	if errorEdit != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "update error " + errorEdit.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data updated",
		"data":    data,
	})
}

func showStatsShortener(c *fiber.Ctx) error {
	url := c.Params("url")
	// call grpc client service

	newContext := middleware.CreateGrpcContext(c.Context(), middleware.GetToken(c))
	data, errorGetStatistic := shortenerClientService.Client.Statistic(
		newContext,
		&pbuild.GetStatisticByUrlRequest{
			BaseUrl: url,
			UserId:  getUserId(c),
		},
	)
	if errorGetStatistic != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "error " + errorGetStatistic.Error(),
		})
	}
	return c.Status(fiber.StatusOK).JSON(model.GetStatisticResponse{
		Day:     data.Day,
		Week:    data.Week,
		Month:   data.Month,
		AllTime: data.AllTime,
	})
}
