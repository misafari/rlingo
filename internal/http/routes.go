package http

import (
	"database/sql"

	"github.com/labstack/echo/v4"
	"github.com/misafari/rlingo/internal/domain/translation"
	"github.com/misafari/rlingo/internal/http/handlers"
	"github.com/misafari/rlingo/internal/infra/db"
)

func RegisterRoutes(e *echo.Echo, dbConn *sql.DB) {
	repo := db.NewTranslationRepo(dbConn)
	service := translation.NewService(repo)
	handler := handlers.NewTranslationHandler(service)

	g := e.Group("/api/v1/translations")
	g.POST("", handler.Create)
	g.GET("/:tenant_id/:project_id", handler.ListByProject)
	g.GET("/:tenant_id/:id", handler.Get)
	g.PUT("/:tenant_id/:id", handler.Update)
	g.PUT("/:tenant_id/:id/approve", handler.Approve)
	g.DELETE("/:tenant_id/:id", handler.Delete)

	userRepo := db.NewUserRepo(dbConn)
    authHandler := handlers.NewAuthHandler(userRepo)

	authGroup := e.Group("/api/v1/auth")
	authGroup.POST("/signup", authHandler.Signup)
	authGroup.POST("/login", authHandler.Login)
}
