package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/misafari/rlingo/internal/delivery/http/handler"
	"github.com/misafari/rlingo/internal/infrastructure/database"
	"github.com/misafari/rlingo/internal/repository/postgres"
	"github.com/misafari/rlingo/internal/usecase/translation"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConf := database.Config{
		Host:           "localhost",
		Port:           5455,
		User:           "admin",
		Password:       "admin",
		Database:       "rlingo",
		MaxConnections: 25,
		MinConnections: 5,
	}

	postgresqlDatabaseConnectionPool, err := database.Connect(ctx, dbConf)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	defer database.Close()

	translationRepository := postgres.NewTranslationRepository(postgresqlDatabaseConnectionPool)

	translationModifyingUseCase := translation.NewModifyingUseCase(translationRepository)

	app := fiber.New(fiber.Config{
		AppName: "Rlingo API v1.0",
	})

	app.Use(logger.New())

	api := app.Group("/api/v1")

	translationHttpHandler := handler.NewTranslationHandler(translationModifyingUseCase)
	translationApiGroup := api.Group("/translations")
	translationApiGroup.Post("/", translationHttpHandler.CreateTranslation)

	if err := app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
