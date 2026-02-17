package middlewares

import (
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	jwtware "github.com/saveblush/gofiber3-contrib/jwt"
)

func Protected() fiber.Handler {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET_KEY"))},

		ErrorHandler: func(c fiber.Ctx, err error) error {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid or expired JWT",
			})
		},
	})
}

func PermissionAdmin(c fiber.Ctx) error {
	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	if claims["role"] != "admin" {
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"message": "Permission denied: admin role required",
		})
	}

	return c.Next()
}
