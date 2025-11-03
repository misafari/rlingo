package main

import (
    "database/sql"
    "log"

    "github.com/labstack/echo/v4"
    _ "github.com/lib/pq"
    "tms/internal/http"
)

func main() {
    db, err := sql.Open("postgres", "postgres://tms:tms@localhost:5432/tms?sslmode=disable")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    e := echo.New()
    http.RegisterRoutes(e, db)

    log.Println("Server running on :8080")
    e.Logger.Fatal(e.Start(":8080"))
}
