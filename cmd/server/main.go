package main

import (
	"log"

	dbsqlc "github.com/RahulPoluru01/user-service/db/sqlc"
	"github.com/RahulPoluru01/user-service/internal/handler"
	"github.com/RahulPoluru01/user-service/internal/logger"
	"github.com/RahulPoluru01/user-service/internal/middleware"
	"github.com/RahulPoluru01/user-service/internal/repository"
	"github.com/RahulPoluru01/user-service/internal/routes"
	"github.com/RahulPoluru01/user-service/internal/service"
	"github.com/gofiber/fiber/v2"
)

func main() {
	logger.Init()
	logger.Log.Info("server starting")

	app := fiber.New()

	app.Use(middleware.RequestID())
	app.Use(middleware.RequestLogger())

	dsn := "postgres://postgres:postgres@localhost:5433/user_service?sslmode=disable"

	db, err := repository.NewDB(dsn)
	if err != nil {
		log.Fatal(err)
	}

	queries := dbsqlc.New(db)

	userRepo := repository.NewPostgresUserRepository(queries)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	routes.Setup(app, userHandler)

	log.Fatal(app.Listen(":3000"))
}
