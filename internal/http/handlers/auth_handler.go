package handlers

import (
    "net/http"
    "time"

    "github.com/golang-jwt/jwt/v5"
    "github.com/google/uuid"
    "github.com/labstack/echo/v4"
    "github.com/misafari/rlingo/internal/domain/user"
    "github.com/misafari/rlingo/internal/security"
)

var jwtSecret = []byte("super-secret-key")

type AuthHandler struct {
    userRepo user.Repository
}

func NewAuthHandler(repo user.Repository) *AuthHandler {
    return &AuthHandler{userRepo: repo}
}

type SignupRequest struct {
    TenantID string `json:"tenant_id"`
    Name     string `json:"name"`
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) Signup(c echo.Context) error {
    var req SignupRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid input")
    }

    hash, err := security.HashPassword(req.Password)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "hashing failed")
    }

    tenantID, _ := uuid.Parse(req.TenantID)
    u := &user.User{
        ID:           uuid.New(),
        TenantID:     tenantID,
        Email:        req.Email,
        PasswordHash: hash,
        Name:         req.Name,
        Role:         "user",
    }

    if err := h.userRepo.Create(c.Request().Context(), u); err != nil {
        return echo.NewHTTPError(http.StatusConflict, "email already exists")
    }

    return c.JSON(http.StatusCreated, map[string]string{"message": "user created"})
}

type LoginRequest struct {
    Email    string `json:"email"`
    Password string `json:"password"`
}

func (h *AuthHandler) Login(c echo.Context) error {
    var req LoginRequest
    if err := c.Bind(&req); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, "invalid payload")
    }

    u, err := h.userRepo.FindByEmail(c.Request().Context(), req.Email)
    if err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
    }

    if err := security.VerifyPassword(u.PasswordHash, req.Password); err != nil {
        return echo.NewHTTPError(http.StatusUnauthorized, "invalid credentials")
    }

    claims := jwt.MapClaims{
        "sub":   u.ID.String(),
        "email": u.Email,
        "tenant": u.TenantID.String(),
        "role":  u.Role,
        "exp":   time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    signed, _ := token.SignedString(jwtSecret)

    return c.JSON(http.StatusOK, map[string]string{"token": signed})
}
