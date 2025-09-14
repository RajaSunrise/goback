package middleware

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// RequestLogger logs information about each request
func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		err := c.Next()
		stop := time.Now()

		log.Printf(
			"[%s] %s %s - %v %s",
			c.IP(),
			c.Method(),
			c.Path(),
			stop.Sub(start),
			c.GetRespHeader("X-Request-ID"),
		)

		return err
	}
}
