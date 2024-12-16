package grpc

import (
	"context"
	"fmt"
	api "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/pkg/api/v1"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	grpcServer *grpc.Server
	listener   net.Listener
}

func New(ctx context.Context, port int) (*Server, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	api.RegisterServiceDeployerServer(grpcServer, nil) // edit!!!

	return &Server{grpcServer, lis}, nil
}

func (s *Server) Start(ctx context.Context) error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		// logger.GetLoggerFromCtx(ctx).Info(ctx, "starting gRPC server", zap.Int("port", s.listener.Addr().(*net.TCPAddr).Port))
		return s.grpcServer.Serve(s.listener)
	})

	return eg.Wait()
}

func (s *Server) Stop(ctx context.Context) {

	s.grpcServer.GracefulStop()
	//l := logger.GetLoggerFromCtx(ctx)
	//if l != nil {
	//	l.Info(ctx, "gRPC server stopped")
	//}
}
