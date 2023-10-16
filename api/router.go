package api

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"jemmy-sapta/configs"
	"jemmy-sapta/middleware"
)

func InitialRoute(config configs.AppConfig) {
	router := fiber.New()
	router.Use(
		cors.New(cors.Config{
			AllowOrigins: "*",
			AllowHeaders: "Origin, Content-Type, Accept",
		}),
	)

	router.Get("/", wellcome)
	router.Get("/:url", redirectUrl)

	router.Post("/register", register)
	router.Post("/login", login)

	router.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(config.SECRET_KEY)},
	}))

	router.Post("/create", middleware.IsValidAuthToken, createUrlShortener)
	router.Put("/update", middleware.IsValidAuthToken, updateUrlShortener)
	router.Get("/stats/:url", middleware.IsValidAuthToken, showStatsShortener)

	router.Listen(config.APP_HOST)
}
