package http

import (
    "database/sql"

    "github.com/labstack/echo/v4"
    "tms/internal/domain/translation"
    "tms/internal/infra/db"
    "tms/internal/http/handlers"
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
}
