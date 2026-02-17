package middlewares

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
)

func CheckMiddle(c fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	log.Printf(
		"URL = %s, Method = %s, Status = %d, Duration = %s, IP = %s",
		c.OriginalURL(), c.Method(), c.Response().StatusCode(), time.Since(start), c.IP(),
	)

	return err
}
