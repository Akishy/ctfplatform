package main

import (
	"context"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/application"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

// для логера
const (
	serviceName = "Control"
)

func main() {
	ctx := context.Background()

	mainLogger := logger.New(serviceName)
	ctx = context.WithValue(ctx, logger.LoggerKey, mainLogger)

	cfg, err := config.New()
	if err != nil {
		panic(fmt.Errorf("failed to load config: %w", err))
	}

	app, err := application.NewApplication(ctx, cfg)

	if err != nil {
		panic(fmt.Errorf("failed to create application: %w", err))
	}

	graceCh := make(chan os.Signal, 1)
	signal.Notify(graceCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := app.Run(); err != nil {
			mainLogger.Error(ctx, err.Error())
		}
	}()

	<-graceCh

	if err := app.Shutdown(ctx); err != nil {
		mainLogger.Error(ctx, err.Error())
	}

	mainLogger.Info(ctx, "Server stopped")
}
