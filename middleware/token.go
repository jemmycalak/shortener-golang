package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"strings"
	"time"
)

func CreateAuthToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	tokenResult, errorSignedToken := token.SignedString([]byte(appConfig.SECRET_KEY))
	if errorSignedToken != nil {
		return "", errorSignedToken
	}
	return tokenResult, nil
}

func CreateClaims(userId int64) jwt.MapClaims {
	return jwt.MapClaims{
		"userId": userId,
		"admin":  true,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	}
}

func GetTokenClaims(c *fiber.Ctx) jwt.MapClaims {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	return claims
}

func GetToken(c *fiber.Ctx) string {
	var accessToken string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		accessToken = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("token") != "" {
		accessToken = c.Cookies("token")
	}

	return accessToken
}
