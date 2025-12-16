package main

import (
	"github.com/RahulPoluru01/user-service/internal/logger"
	"github.com/RahulPoluru01/user-service/internal/middleware"
	"github.com/RahulPoluru01/user-service/internal/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.Init()
	logger.Log.Info("server starting")

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	routes.Setup(app)

	app.Listen(":3000")
}
