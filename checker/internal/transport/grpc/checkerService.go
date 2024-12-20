package grpc

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/hashUtils"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/checkerService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/vulnServiceService"
	proto "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/pkg/api/v1"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CheckerService struct {
	proto.UnimplementedCheckerSystemServer
	checkerService *checkerService.Service
	vulnService    *vulnServiceService.Service
	logger         *zap.Logger
}

func NewCheckerService(service *checkerService.Service, vulnServiceService *vulnServiceService.Service, logger *zap.Logger) *CheckerService {
	return &CheckerService{
		checkerService: service,
		vulnService:    vulnServiceService,
		logger:         logger,
	}
}

func (s *CheckerService) RegisterChecker(
	_ context.Context,
	req *proto.RegisterCheckerRequest) (*proto.RegisterCheckerResponse, error) {

	codeArchive := req.CodeArchive
	response := &proto.RegisterCheckerResponse{
		Id: "",
	}
	_, err := s.checkerService.CheckerImgService.GetByRaw(codeArchive)
	if err == nil {
		s.logger.Info("checker image already exists", zap.String("code_archive", codeArchive))
		//response.Id = possibleCheckerImg.Uuid.String()
		return response, status.Error(codes.Canceled, "checker image already exists")
	}

	codeArchiveHash, err := hashUtils.GenerateImgHash(codeArchive)
	if err != nil {
		s.logger.Error("failed to generate code archive hash", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to generate code archive hash")
	}
	s.logger.Debug("code archive hash", zap.String("code_archive", codeArchiveHash))
	checkerImg := checkerImgDomain.CheckerImg{
		Uuid:        uuid.New(),
		Hash:        codeArchiveHash,
		CodeArchive: codeArchive,
	}
	if err := s.checkerService.CheckerImgService.Create(&checkerImg); err != nil {
		s.logger.Error("failed to create checker image", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create checker image")
	}
	s.logger.Debug("checker image successfully created", zap.String("checkerImg", fmt.Sprintf("%v", checkerImg)))

	newChecker := checkerDomain.Checker{
		UUID:       uuid.New(),
		CheckerImg: &checkerImg,
		Ip:         "",
		WebPort:    0,
	}
	if err := s.checkerService.Create(&newChecker); err != nil {
		s.logger.Error("failed to create checker", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create checker")
	}
	s.logger.Debug("checker successfully created", zap.String("checker", fmt.Sprintf("%v", newChecker)))
	response.Id = newChecker.UUID.String()

	return response, nil
}

func (s *CheckerService) PingChecker(
	_ context.Context,
	req *proto.PingCheckerRequest) (*proto.PingCheckerResponse, error) {
	checkerUuid, err := uuid.Parse(req.Id)
	if err != nil {
		s.logger.Error("failed to parse checker uuid", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse req id uuid")
	}

	vulnServiceList, err := s.vulnService.GetActiveList(checkerUuid)
	if err != nil {
		s.logger.Error("failed to get vulnService list", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get vuln services")
	}

	vulnServiceInfos := make([]*proto.VulnServiceInfo, 0)

	for _, vulnService := range vulnServiceList {
		vulnServiceInfo := &proto.VulnServiceInfo{
			StatusCode: int32(vulnService.StatusCode),
			Message:    vulnService.Message,
			WebPort:    int32(vulnService.WebPort),
			Ip:         vulnService.Ip,
			InstanceId: vulnService.Uuid.String(),
			LastCheck:  vulnService.LastCheck,
		}
		vulnServiceInfos = append(vulnServiceInfos, vulnServiceInfo)
	}

	response := &proto.PingCheckerResponse{
		Message:           "",
		VulnServicesInfos: vulnServiceInfos,
	}

	return response, nil

}

func (s *CheckerService) CreateVulnService(_ context.Context, req *proto.CreateVulnServiceRequest) (*proto.CreateVulnServiceResponse, error) {
	checkerUUID, err := uuid.Parse(req.GetServiceId())
	if err != nil {
		s.logger.Error("failed to parse req id uuid", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse req id uuid")
	}
	if exists := s.checkerService.Exists(checkerUUID); !exists {
		s.logger.Error("failed to find checker uuid", zap.String("req id", req.GetServiceId()))
		return nil, status.Errorf(codes.InvalidArgument, "failed to find checker uuid")
	}

	vulnService := vulnServiceDomain.VulnService{
		Uuid:       uuid.New(),
		Ip:         req.GetIp(),
		WebPort:    int(req.GetWebPort()),
		StatusCode: vulnServiceDomain.DOWN,
		Message:    "",
		CheckerId:  checkerUUID,
		LastCheck:  -1,
		Active:     true,
	}

	if err := s.vulnService.Create(&vulnService); err != nil {
		s.logger.Error("failed to create vulnService", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create vulnService")
	}
	return &proto.CreateVulnServiceResponse{
		InstanceId: vulnService.Uuid.String(),
		Success:    true,
	}, nil
}

func (s *CheckerService) StopVulnServices(_ context.Context, req *proto.StopVulnServicesRequest) (*proto.StopVulnServicesResponse, error) {
	for _, id := range req.GetInstanceIds() {
		vulnServiceUuid, err := uuid.Parse(id)
		if err != nil {
			s.logger.Error("failed to parse req id uuid", zap.Error(err))
			return nil, status.Errorf(codes.InvalidArgument, "failed to parse req id uuid")
		}
		err = s.vulnService.DeactivateVulnService(vulnServiceUuid)
		if err != nil {
			s.logger.Error("failed to deactivate vulnService", zap.Error(err))
			return nil, status.Errorf(codes.Internal, "failed to deactivate vulnService")
		}
	}

	return &proto.StopVulnServicesResponse{
		Success: true,
	}, nil
}
func (s *CheckerService) PingVulnService(_ context.Context, req *proto.PingVulnServiceRequest) (*proto.VulnServiceInfo, error) {
	vulnServiceUuid, err := uuid.Parse(req.InstanceId)
	if err != nil {
		s.logger.Error("failed to parse req id uuid", zap.Error(err))
		return nil, status.Errorf(codes.InvalidArgument, "failed to parse req id uuid")
	}

	vulnService, err := s.vulnService.Get(vulnServiceUuid)
	if err != nil {
		s.logger.Error("failed to get vulnService info", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to get vulnService info")
	}

	response := &proto.VulnServiceInfo{
		StatusCode: int32(vulnService.StatusCode),
		Message:    vulnService.Message,
		WebPort:    int32(vulnService.WebPort),
		Ip:         vulnService.Ip,
		InstanceId: vulnService.Uuid.String(),
		LastCheck:  vulnService.LastCheck,
	}

	return response, nil
}
