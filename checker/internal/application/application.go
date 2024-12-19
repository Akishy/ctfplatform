package application

import (
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/adapters/memory"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/logger"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/checkerImgService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/checkerService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/flagGeneratorService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/vulnServiceService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/transport/grpc"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/transport/http"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"time"
)

const defaultCheckDelay = 5 * time.Second

func Init() {
	config.InitConfig()
	lgr, err := logger.NewLogger()
	if err != nil {
		panic(err)
	}
	storage := memory.NewStorage()

	flagGenService := flagGeneratorService.NewService(storage)
	checkImgService := checkerImgService.NewService(storage, lgr)
	checkService := checkerService.NewService(storage, flagGenService, checkImgService, lgr)
	vulnServiceServ := vulnServiceService.NewService(storage)
	grpcCheckService := grpc.NewCheckerService(checkService, vulnServiceServ, lgr)
	grpcServer := grpc.New(lgr, grpcCheckService)

	eg := errgroup.Group{}

	eg.Go(func() error {
		return grpcServer.Start()
	})

	eg.Go(func() error {
		httpServer := http.New(4010, vulnServiceServ, checkService, lgr)
		return httpServer.Start()
	})

	go func() {
		fmt.Println("checkerGoroutine started")
		time.Sleep(4 * defaultCheckDelay)
		for {
			fmt.Println("I sleep")
			time.Sleep(defaultCheckDelay)
			_ = checkService.Check()
		}
	}()

	panic(eg.Wait())
}

func InitFx() {
	config.InitConfig()
	app := fx.New(
		logger.FxOption,
		grpc.ServerFxOption,
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger { return &fxevent.ZapLogger{Logger: log} }),
	)
	app.Run()
}
