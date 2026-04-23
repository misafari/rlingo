package main

import (
	"context"
	"log/slog"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiblogger "github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
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
	log := newLogger(os.Getenv("APP_ENV"))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	dbConf := database.Config{
		Host:           "100.94.100.79",
		Port:           5433,
		User:           "admin",
		Password:       "admin",
		Database:       "rlingo",
		MaxConnections: 25,
		MinConnections: 5,
	}

	postgresqlDatabaseConnectionPool, err := database.Connect(ctx, dbConf)
	if err != nil {
		log.Error("Error connecting to database: %v", err)
	}

	defer database.Close()
	defer postgresqlDatabaseConnectionPool.Close()

	queries := generatedQueries.New(postgresqlDatabaseConnectionPool)

	projectRepository := projectPostgresqlRepositoryImpl.NewRepository(queries, postgresqlDatabaseConnectionPool)
	projectService := project.NewService(projectRepository)

	ur := identityRepo.NewUserRepositoryPostgresImpl(queries, postgresqlDatabaseConnectionPool)
	ts := identity.NewTokenService()
	us := identity.NewUserService(ur, ts)

	tenantRepo := identityRepo.NewTenantRepositoryPostgresImpl(queries, postgresqlDatabaseConnectionPool)
	tenantSvc := identity.NewTenantService(tenantRepo)

	app := fiber.New(fiber.Config{
		ErrorHandler: middleware.ErrorHandler,
		AppName:      "Rlingo API v1.0",
	})

	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(fiblogger.New(fiblogger.Config{
		Format: "${time} | ${method} ${path} | ${status} | ${latency}\n",
	}))

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "GET, POST, PUT, PATCH, DELETE, OPTIONS",
	}))

	v1 := app.Group("/api/v1")

	authMiddleware := middleware.NewAuthMiddleware(ts, tenantSvc, log)

	secureApi := v1.Group(
		"/secure",
		authMiddleware.TokenValidationFilter(),
		authMiddleware.TenantFilter(),
	)
	unsecureApi := v1.Group("/unsecure")

	phh := projectHttpHandler.NewHttpHandler(projectService)
	papig := secureApi.Group("/projects")
	papig.Get("/", phh.GetAll)
	papig.Post("/", phh.Create)
	papig.Delete("/:id", phh.DeleteOneById)
	papig.Put("/:id", phh.Update)

	iuhh := identityHttpHandler.NewUserHttpHandler(us)
	iuapig := unsecureApi.Group("/auth")
	iuapig.Post("/sign-up", iuhh.Signup)
	iuapig.Post("/sign-in", iuhh.Signin)

	if err = app.Listen(":8000"); err != nil {
		log.Error(err.Error())
	}
}

func newLogger(env string) *slog.Logger {
	if env == "production" {
		return slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}
	return slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
}
