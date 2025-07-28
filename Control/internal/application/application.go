package application

import (
	"context"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/repository"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/transport/rest"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/db"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
)

type Application struct {
	context    context.Context
	server     *rest.Server
	repository rest.Repository
	config     *config.Config
}

func NewApplication(ctx context.Context, config *config.Config) (*Application, error) {
	pg, err := db.NewDB(ctx, &config.PostgresConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres: %v", err)
	}
	// инициализация репозитория
	// каст в самый большой интерфейс (такой хранит сервер) по указателю
	var repo rest.Repository = repository.NewRepository(pg)

	// создание REST сервера
	server := rest.NewServer(ctx, config, repo)

	// создание экземпляра приложения
	app := &Application{
		context:    ctx,
		server:     server,
		repository: repo,
		config:     config,
	}

	return app, nil
}

// Run запускает приложение
func (app *Application) Run() error {
	return app.server.Start()
}

// Shutdown gracefully останавливает приложение
func (app *Application) Shutdown(ctx context.Context) error {
	l := logger.GetLoggerFromCtx(ctx)
	l.Info(app.context, "stopping Control-server...")

	if err := app.server.GracefulStop(); err != nil {
		return err
	}

	if err := app.repository.Close(); err != nil {
		return err
	}

	return nil
}
