package server

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/internal/config"
	pb "gitlab.crja72.ru/gospec/go4/ctfplatform/deployer/pkg/proto"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"sync"
)

// Server структура веб-сервера.
type Server struct {
	pb.UnimplementedServiceDeployerServer
	logger    *zap.Logger
	services  map[string]*Service
	instances map[string]*Instance
	mutex     sync.RWMutex
}

type Service struct {
	ID          string
	CodeArchive string
}

type Instance struct {
	ID         string
	ServiceID  string
	SSHPort    int32
	WebPort    int32
	IP         string
	StatusCode int32
	Message    string
}

func NewServer(log *zap.Logger) *Server {
	return &Server{
		logger:    log,
		services:  make(map[string]*Service),
		instances: make(map[string]*Instance),
	}
}

func (s *Server) CreateService(ctx context.Context, req *pb.RegisterServiceRequest) (*pb.RegisterServiceResponse, error) {
	if req.CodeArchive == "" {
		return nil, errors.New("code_archive cannot be empty")
	}

	serviceID := uuid.New().String()

	s.mutex.Lock()
	s.services[serviceID] = &Service{
		ID:          serviceID,
		CodeArchive: req.CodeArchive,
	}
	s.mutex.Unlock()

	s.logger.Info("Service created", zap.String("id", serviceID))
	return &pb.RegisterServiceResponse{Id: serviceID}, nil
}

func (s *Server) PingService(ctx context.Context, req *pb.PingServiceRequest) (*pb.PingServiceResponse, error) {
	s.mutex.RLock()
	service, exists := s.services[req.Id]
	if !exists {
		s.mutex.RUnlock()
		return nil, errors.New("service not found")
	}
	var instances []*pb.InstanceInfo
	for _, instance := range s.instances {
		if instance.ServiceID == service.ID {
			instances = append(instances, &pb.InstanceInfo{
				StatusCode: instance.StatusCode,
				Message:    instance.Message,
				SshPort:    instance.SSHPort,
				WebPort:    instance.WebPort,
				Ip:         instance.IP,
				InstanceId: instance.ID,
			})
		}
	}
	s.mutex.RUnlock()

	s.logger.Info("PingService", zap.String("service_id", req.Id))
	return &pb.PingServiceResponse{
		Message:   "Service found",
		Instances: instances,
	}, nil
}

func (s *Server) PingInstance(ctx context.Context, req *pb.PingInstanceRequest) (*pb.InstanceInfo, error) {
	s.mutex.RLock()
	instance, exists := s.instances[req.InstanceId]
	s.mutex.RUnlock()
	if !exists {
		return nil, errors.New("instance not found")
	}

	s.logger.Info("PingInstance", zap.String("instance_id", req.InstanceId))
	return &pb.InstanceInfo{
		StatusCode: instance.StatusCode,
		Message:    instance.Message,
		SshPort:    instance.SSHPort,
		WebPort:    instance.WebPort,
		Ip:         instance.IP,
		InstanceId: instance.ID,
	}, nil
}

func (s *Server) StartInstances(ctx context.Context, req *pb.StartInstancesRequest) (*pb.StartInstancesResponse, error) {
	if req.Count <= 0 {
		return nil, errors.New("count must be greater than 0")
	}
	s.mutex.RLock()
	_, exists := s.services[req.ServiceId]
	s.mutex.RUnlock()
	if !exists {
		return nil, errors.New("service not found")
	}

	var instanceIDs []string
	for i := 0; i < int(req.Count); i++ {
		instanceID := uuid.New().String()
		instance := &Instance{
			ID:         instanceID,
			ServiceID:  req.ServiceId,
			SSHPort:    22,
			WebPort:    8080,
			IP:         "127.0.0.1",
			StatusCode: 200,
			Message:    "Running",
		}
		instanceIDs = append(instanceIDs, instanceID)

		s.mutex.Lock()
		s.instances[instanceID] = instance
		s.mutex.Unlock()
	}

	s.logger.Info("StartInstances", zap.String("service_id", req.ServiceId), zap.Int("count", len(instanceIDs)))
	return &pb.StartInstancesResponse{
		InstanceIds: instanceIDs,
		Success:     true,
	}, nil
}

func (s *Server) StopInstances(ctx context.Context, req *pb.StopInstancesRequest) (*pb.StopInstancesResponse, error) {
	s.mutex.Lock()
	for _, id := range req.InstanceIds {
		delete(s.instances, id)
	}
	s.mutex.Unlock()

	s.logger.Info("StopInstances", zap.Int("count", len(req.InstanceIds)))
	return &pb.StopInstancesResponse{Success: true}, nil
}

// StartServer запускает сервер и добавляет хук для остановки.
// StartServer запускает gRPC-сервер и добавляет хук для остановки.
func StartServer(lc fx.Lifecycle, server *Server, cfg *config.Config) {
	grpcServer := grpc.NewServer()
	pb.RegisterServiceDeployerServer(grpcServer, server)

	listener, err := net.Listen("tcp", cfg.AppPort)
	if err != nil {
		server.logger.Fatal("Failed to start listener", zap.Error(err))
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				server.logger.Info("Starting gRPC server", zap.String("addr", cfg.AppPort))
				if err := grpcServer.Serve(listener); err != nil {
					server.logger.Fatal("gRPC server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.logger.Info("Stopping gRPC server...")
			grpcServer.GracefulStop()
			return nil
		},
	})
}

// Module предоставляет зависимости сервера.
var Module = fx.Options(
	fx.Provide(NewServer),
	fx.Invoke(StartServer),
)
