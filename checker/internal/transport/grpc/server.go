package grpc

import (
	"context"
	api "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/pkg/api/v1"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	logger     *zap.Logger
}

var ServerFxOption fx.Option = fx.Options(fx.Provide(NewFx), fx.Invoke(func(*Server) {}))

func New(logger *zap.Logger, checkerService *CheckerService) *Server {
	grpcServer := grpc.NewServer()
	api.RegisterCheckerSystemServer(grpcServer, checkerService)
	return &Server{
		grpcServer: grpcServer,
		logger:     logger,
	}
}

func NewFx(lc fx.Lifecycle, logger *zap.Logger, checkerService *CheckerService) *Server {
	grpcServer := grpc.NewServer()
	api.RegisterCheckerSystemServer(grpcServer, checkerService)

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			logger.Info("Starting gRPC server")
			lis, err := net.Listen("tcp", ":4000")
			if err != nil {
				logger.Error("gRPC server initialization failed", zap.Error(err))
			}
			go func() {
				if err := grpcServer.Serve(lis); err != nil {
					logger.Error("gRPC server initialization failed", zap.Error(err))
				}
			}()

			return nil
		},
		OnStop: func(context.Context) error {
			logger.Info("Stopping gRPC server")
			grpcServer.GracefulStop()
			return nil
		},
	})

	return &Server{
		grpcServer: grpcServer,
		logger:     logger,
	}
}

func (s *Server) Start() error {
	s.logger.Info("Starting gRPC server")
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		s.logger.Error("grpc server initialization failed", zap.Error(err))
	}
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
