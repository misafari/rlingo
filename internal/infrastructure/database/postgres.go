package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

type Config struct {
	Host           string
	Port           uint16
	User           string
	Password       string
	Database       string
	MaxConnections int32
	MinConnections int32
}

func Connect(ctx context.Context, cfg Config) (*pgxpool.Pool, error) {
	connectionStr := fmt.Sprintf(
		"postgresql://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)

	return ConnectWithConnectionStr(ctx, connectionStr, &cfg.MaxConnections, &cfg.MinConnections)
}

func ConnectWithConnectionStr(ctx context.Context, connectionStr string, maxConnection *int32, minConnection *int32) (*pgxpool.Pool, error) {
	poolConfig, err := pgxpool.ParseConfig(connectionStr)

	if err != nil {
		log.Fatalf("Unable to parse config: %v\n", err)
	}

	if maxConnection != nil && *maxConnection > 0 {
		poolConfig.MaxConns = *maxConnection
	} else {
		poolConfig.MaxConns = 25
	}

	if minConnection != nil && *minConnection > 0 {
		poolConfig.MinConns = *minConnection
	} else {
		poolConfig.MinConns = 5
	}

	poolConfig.MaxConnLifetime = 1 * time.Hour
	poolConfig.MaxConnIdleTime = 30 * time.Minute
	poolConfig.HealthCheckPeriod = 1 * time.Minute

	Pool, err = pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	if err := Pool.Ping(ctx); err != nil {
		log.Fatalf("Unable to ping database: %v\n", err)
	}

	log.Println("Connected to database successfully")
	log.Printf("Max Connection: %d", poolConfig.MaxConns)
	log.Printf("Min Connection: %d", poolConfig.MinConns)

	return Pool, nil
}

func Close() {
	if Pool != nil {
		Pool.Close()
		log.Println("Database connection pool closed")
	} else {
		log.Println("No database connection pool to close")
	}
}
