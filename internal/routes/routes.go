package routes

import (
	"github.com/RahulPoluru01/user-service/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App, h *handler.UserHandler) {
	app.Get("/health", handler.Health)

	app.Post("/users", h.CreateUser)
	app.Get("/users/:id", h.GetUserByID)
	app.Get("/users", h.ListUsers)
	app.Put("/users/:id", h.UpdateUser)
	app.Delete("/users/:id", h.DeleteUser)
}
