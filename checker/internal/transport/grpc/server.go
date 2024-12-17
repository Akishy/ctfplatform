package grpc

import (
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

var ServerFxOption fx.Option = fx.Options(fx.Provide(New), fx.Invoke(func(*Server) {}))

func New(logger *zap.Logger, checkerService *CheckerService) *Server {
	grpcServer := grpc.NewServer()
	api.RegisterCheckerSystemServer(grpcServer, checkerService)

	return &Server{grpcServer, logger}
}

func (s *Server) Start() error {
	lis, err := net.Listen("tcp", ":4000")
	if err != nil {
		s.logger.Error("grpc server initialization failed", zap.Error(err))
	}
	return s.grpcServer.Serve(lis)
}

func (s *Server) Stop() {
	s.grpcServer.GracefulStop()
}
