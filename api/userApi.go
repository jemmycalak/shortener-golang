package api

import (
	"github.com/gofiber/fiber/v2"
	"jemmy-sapta/client/service"
	"jemmy-sapta/middleware"
	"jemmy-sapta/model"
	"jemmy-sapta/protos/pbuild"
)

var userClientService service.UserClientService

func InitUserApi(
	clientService service.UserClientService,
) {
	userClientService = clientService
}

func register(c *fiber.Ctx) error {
	var body model.User
	if errorParsing := c.BodyParser(&body); errorParsing != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing " + errorParsing.Error(),
		})
	}
	responeRegister, errorRegister := userClientService.Client.CreateUser(
		c.Context(),
		&pbuild.CreateUserRequest{
			Username: body.Username,
			Password: body.Password,
		})
	if errorRegister != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error create" + errorRegister.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(responeRegister)
}

func login(c *fiber.Ctx) error {
	var body model.User
	if errorParsing := c.BodyParser(&body); errorParsing != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error parsing " + errorParsing.Error(),
		})
	}

	responeLogin, errorLogin := userClientService.Client.GetUserBy(
		c.Context(),
		&pbuild.GetUserByRequest{
			Username: body.Username,
			Password: body.Password,
		},
	)
	if errorLogin != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error login" + errorLogin.Error(),
		})
	}

	claims := middleware.CreateClaims(responeLogin.Id)
	token, errorCreateToken := middleware.CreateAuthToken(claims)
	if errorCreateToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error create token" + errorCreateToken.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}
