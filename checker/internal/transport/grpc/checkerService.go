package grpc

import (
	checkerService2 "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/checkerService"
	proto "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/pkg/api/v1"
)

type CheckerService struct {
	proto.UnimplementedServiceDeployerServer
	checkerService *checkerService2.Service
}

func NewCheckerService(service *checkerService2.Service) *CheckerService {
	return &CheckerService{checkerService: service}
}
