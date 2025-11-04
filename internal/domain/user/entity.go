package user

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID           uuid.UUID `json:"id"`
    TenantID     uuid.UUID `json:"tenant_id"`
    Email        string    `json:"email"`
    PasswordHash string    `json:"-"`
    Name         string    `json:"name"`
    Role         string    `json:"role"`
    CreatedAt    time.Time `json:"created_at"`
    UpdatedAt    time.Time `json:"updated_at"`
}
