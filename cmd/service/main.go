package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fisdemire/pickup_point_service/internal/auth"
	"github.com/fisdemire/pickup_point_service/internal/config"
	"github.com/fisdemire/pickup_point_service/internal/service"
	"github.com/fisdemire/pickup_point_service/internal/storage/postgres"
	"github.com/fisdemire/pickup_point_service/internal/transport/httpapi"
)

func main() {
	if err := run(); err != nil {
		slog.Error("fatal", "err", err)
		os.Exit(1)
	}
}

func run() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg, err := config.Load()
	if err != nil {
		return err
	}
	if err := postgres.Migrate(ctx, cfg.DatabaseURL); err != nil {
		return err
	}
	pool, err := postgres.NewPool(ctx, cfg.DatabaseURL)
	if err != nil {
		return err
	}
	defer pool.Close()

	store := postgres.NewStore(pool)
	tokens := auth.NewTokenManager(cfg.JWTSecret, cfg.JWTTTL)

	srv := &http.Server{
		Addr:              cfg.HTTPAddr,
		Handler:           httpapi.NewRouter(httpapi.Deps{Auth: service.NewAuth(store, tokens), Tokens: tokens}),
		ReadHeaderTimeout: 5 * time.Second,
	}
	go func() {
		<-ctx.Done()
		sctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		_ = srv.Shutdown(sctx)
	}()
	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}
