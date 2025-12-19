package middleware

import (
	"time"

	"github.com/RahulPoluru01/user-service/internal/logger"
	"github.com/gofiber/fiber/v2"
)

func RequestLogger() fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		err := c.Next()

		duration := time.Since(start)
		reqID := c.Locals("requestId")

		logger.Log.Info(
			"request completed",
		)

		_ = reqID
		_ = duration

		return err
	}
}
