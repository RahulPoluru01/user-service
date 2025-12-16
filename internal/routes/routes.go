package routes

import (
	"github.com/RahulPoluru01/user-service/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Get("/health", handler.Health)
}
