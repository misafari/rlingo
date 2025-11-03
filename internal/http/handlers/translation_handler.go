package handlers

import (
    "net/http"

    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "tms/internal/domain/translation"
)

type TranslationHandler struct {
    service *translation.Service
}

func NewTranslationHandler(service *translation.Service) *TranslationHandler {
    return &TranslationHandler{service: service}
}

// Create new translation
func (h *TranslationHandler) Create(c echo.Context) error {
    var req struct {
        TenantID  string `json:"tenant_id"`
        ProjectID string `json:"project_id"`
        Key       string `json:"key"`
        Locale    string `json:"locale"`
        Text      string `json:"text"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    tenantID, _ := uuid.Parse(req.TenantID)
    projectID, _ := uuid.Parse(req.ProjectID)
    t, err := h.service.Create(c.Request().Context(), tenantID, projectID, req.Key, req.Locale, req.Text)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusCreated, t)
}

// List all translations by project
func (h *TranslationHandler) ListByProject(c echo.Context) error {
    tenantID := c.Param("tenant_id")
    projectID := c.Param("project_id")

    list, err := h.service.Repo().ListByProject(c.Request().Context(), tenantID, projectID)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.JSON(http.StatusOK, list)
}

// Get single translation by ID
func (h *TranslationHandler) Get(c echo.Context) error {
    tenantID := c.Param("tenant_id")
    id := c.Param("id")

    t, err := h.service.Repo().GetByID(c.Request().Context(), tenantID, id)
    if err != nil {
        return echo.NewHTTPError(http.StatusNotFound, err.Error())
    }
    return c.JSON(http.StatusOK, t)
}

// Update translation text
func (h *TranslationHandler) Update(c echo.Context) error {
    tenantIDStr := c.Param("tenant_id")
    id := c.Param("id")

    var req struct {
        Text string `json:"text"`
    }
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }

    tenantID, _ := uuid.Parse(tenantIDStr)
    if err := h.service.UpdateText(c.Request().Context(), tenantID, id, req.Text); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "updated"})
}

// Approve translation
func (h *TranslationHandler) Approve(c echo.Context) error {
    tenantIDStr := c.Param("tenant_id")
    id := c.Param("id")

    tenantID, _ := uuid.Parse(tenantIDStr)
    if err := h.service.Approve(c.Request().Context(), tenantID, id); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return c.JSON(http.StatusOK, map[string]string{"message": "approved"})
}

// Delete translation
func (h *TranslationHandler) Delete(c echo.Context) error {
    tenantID := c.Param("tenant_id")
    id := c.Param("id")

    if err := h.service.Repo().Delete(c.Request().Context(), tenantID, id); err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
    }
    return c.NoContent(http.StatusNoContent)
}
