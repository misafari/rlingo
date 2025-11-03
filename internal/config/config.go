package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Server   ServerConfig   `yaml:"server"`
}

type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	DBName   string `yaml:"dbname"`
	SSLMode  string `yaml:"sslmode"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

func Load(configPath string) (*Config, error) {
	cfg := &Config{
		Database: DatabaseConfig{
			Host:     "localhost",
			Port:     "5432",
			User:     "tms",
			Password: "tms",
			DBName:   "tms",
			SSLMode:  "disable",
		},
		Server: ServerConfig{
			Port: "8080",
		},
	}

	if configPath != "" {
		if err := loadFromFile(configPath, cfg); err != nil {
			return nil, fmt.Errorf("failed to load config file: %w", err)
		}
	}

	overrideFromEnv(cfg)

	return cfg, nil
}

func loadFromFile(path string, cfg *Config) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return err
	}

	return nil
}

func overrideFromEnv(cfg *Config) {
	if host := getEnv("DB_HOST", ""); host != "" {
		cfg.Database.Host = host
	}
	if port := getEnv("DB_PORT", ""); port != "" {
		cfg.Database.Port = port
	}
	if user := getEnv("DB_USER", ""); user != "" {
		cfg.Database.User = user
	}
	if password := getEnv("DB_PASSWORD", ""); password != "" {
		cfg.Database.Password = password
	}
	if dbname := getEnv("DB_NAME", ""); dbname != "" {
		cfg.Database.DBName = dbname
	}
	if sslmode := getEnv("DB_SSLMODE", ""); sslmode != "" {
		cfg.Database.SSLMode = sslmode
	}
	if port := getEnv("SERVER_PORT", ""); port != "" {
		cfg.Server.Port = port
	}
}

func (d DatabaseConfig) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		d.User, d.Password, d.Host, d.Port, d.DBName, d.SSLMode)
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

