package config

import (
	"errors"
	"os"
	"time"
)

type Config struct {
	HTTPAddr    string
	GRPCAddr    string
	MetricsAddr string
	DatabaseURL string
	JWTSecret   string
	JWTTTL      time.Duration
}

func Load() (Config, error) {
	cfg := Config{
		HTTPAddr:    env("HTTP_ADDR", ":8080"),
		GRPCAddr:    env("GRPC_ADDR", ":3000"),
		MetricsAddr: env("METRICS_ADDR", ":9000"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		JWTSecret:   os.Getenv("JWT_SECRET"),
	}
	ttl, err := time.ParseDuration(env("JWT_TTL", "24h"))
	if err != nil {
		return Config{}, err
	}
	cfg.JWTTTL = ttl
	if cfg.DatabaseURL == "" || cfg.JWTSecret == "" {
		return Config{}, errors.New("DATABASE_URL and JWT_SECRET are required")
	}
	return cfg, nil
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
