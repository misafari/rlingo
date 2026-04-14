package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	generatedQueries "github.com/misafari/rlingo/internal/db/generated"
	"github.com/misafari/rlingo/internal/identity"
	identityHttpHandler "github.com/misafari/rlingo/internal/identity/handler"
	identityRepo "github.com/misafari/rlingo/internal/identity/repository"
	"github.com/misafari/rlingo/internal/infrastructure/database"
	"github.com/misafari/rlingo/internal/project"
	projectHttpHandler "github.com/misafari/rlingo/internal/project/handler"
	projectPostgresqlRepositoryImpl "github.com/misafari/rlingo/internal/project/postgresql"
	"github.com/misafari/rlingo/internal/share/middleware"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConf := database.Config{
		Host:           "100.88.246.49",
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
	defer postgresqlDatabaseConnectionPool.Close()

	queries := generatedQueries.New(postgresqlDatabaseConnectionPool)

	projectRepository := projectPostgresqlRepositoryImpl.NewRepository(queries, postgresqlDatabaseConnectionPool)
	projectService := project.NewService(projectRepository)

	ur := identityRepo.NewUserRepositoryPostgresImpl(queries, postgresqlDatabaseConnectionPool)
	ts := identity.NewTokenService()
	us := identity.NewUserService(ur, ts)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		AppName:      "Rlingo API v1.0",
	})

	app.Use(logger.New(), cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	api := app.Group("/api/v1")

	phh := projectHttpHandler.NewHttpHandler(projectService)
	papig := api.Group("/projects")
	papig.Get("/", phh.GetAll)
	papig.Post("/", phh.Create)
	papig.Delete("/:id", phh.DeleteOneById)
	papig.Put("/:id", phh.Update)

	iuhh := identityHttpHandler.NewUserHttpHandler(us)
	iuapig := api.Group("/auth")
	iuapig.Post("/sign-up", iuhh.Signup)
	iuapig.Post("/sign-in", iuhh.Signin)

	if err = app.Listen(":8000"); err != nil {
		log.Fatal(err)
	}
}
