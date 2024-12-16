package application

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/logger"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func Init() {
	config.InitConfig()
	app := fx.New(
		logger.Module,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: log} }),
	)
	app.Run()
}
