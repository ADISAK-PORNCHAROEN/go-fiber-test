package services

import (
	"fiber-test/internal/models"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

var userAdmin = models.User{
	Username: "admin",
	Password: "1234",
}

func LoginService(username string, password string) (string, error) {
	if username != userAdmin.Username || password != userAdmin.Password {
		return "", fiber.ErrUnauthorized
	}

	token := jwt.New(jwt.SigningMethodHS256)

	clamis := token.Claims.(jwt.MapClaims)
	clamis["username"] = username
	clamis["role"] = "admin"
	clamis["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return fiber.ErrBadGateway.Error(), err
	}

	return t, nil
}
