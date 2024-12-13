package main

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/cache"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/db"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/logger"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/server"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.NewConfig,
		),
		logger.Module,
		db.Module,
		cache.Module,
		server.Module,
	)
	app.Run()
}
