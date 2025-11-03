package main

import (
    "database/sql"
    "log"
    "os"

    "github.com/labstack/echo/v4"
    _ "github.com/lib/pq"
    "github.com/misafari/rlingo/internal/config"
    "github.com/misafari/rlingo/internal/http"
)

func main() {
    configPath := getConfigPath()
    cfg, err := config.Load(configPath)
    if err != nil {
        log.Fatal("Failed to load config:", err)
    }

    db, err := sql.Open("postgres", cfg.Database.DSN())
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }
    defer db.Close()

    if err := db.Ping(); err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    e := echo.New()
    http.RegisterRoutes(e, db)

    addr := ":" + cfg.Server.Port
    log.Printf("Server running on %s", addr)
    e.Logger.Fatal(e.Start(addr))
}

func getConfigPath() string {
    if path := os.Getenv("CONFIG_FILE"); path != "" {
        return path
    }
    if _, err := os.Stat("config.yaml"); err == nil {
        return "config.yaml"
    }
    if _, err := os.Stat("config.yml"); err == nil {
        return "config.yml"
    }
    return ""
}
