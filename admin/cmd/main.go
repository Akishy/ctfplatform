package main

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/cache"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/db"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/logger"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/server"
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
