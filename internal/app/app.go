// Package app configures and runs application.
package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"

	"vladmsnk/billing/config"
	v1 "vladmsnk/billing/internal/controller/http/v1"
	"vladmsnk/billing/internal/usecase"
	"vladmsnk/billing/internal/usecase/repo"
	"vladmsnk/billing/pkg/httpserver"
	"vladmsnk/billing/pkg/logger"
	"vladmsnk/billing/pkg/postgres"
)

// Run creates objects via constructors.
func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	// Repository
	pg, err := postgres.New(cfg.PG.URL, postgres.MaxPoolSize(cfg.PG.PoolMax))
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	// Use case
	BillingUseCase := usecase.NewBillingUseCase(
		repo.New(pg),
	)
	AuthUserCase := usecase.NewAuthUseCase(
		repo.New(pg))

	// HTTP Server
	handler := gin.New()
	v1.NewRouter(handler, l, BillingUseCase, AuthUserCase)
	httpServer := httpserver.New(handler, httpserver.Port(cfg.HTTP.Port))

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	// Shutdown
	err = httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
