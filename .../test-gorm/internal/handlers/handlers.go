package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthCheck godoc
// @Summary      Show the status of server.
// @Description  get the status of server.
// @Tags         health
// @Accept       json
// @Produce      json
// @Success      200  {string}  string    "OK"
// @Router       /health [get]
func HealthCheck(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON("OK")
}
