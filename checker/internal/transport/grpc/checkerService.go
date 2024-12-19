package grpc

import (
	"context"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
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
	possibleCheckerImg, err := s.checkerService.CheckerImgService.Get(codeArchive)
	if err == nil {
		s.logger.Info("checker image already exists", zap.String("code_archive", codeArchive))
		response.Id = possibleCheckerImg.Uuid.String()
		return response, nil
	}

	codeArchiveHash, err := hashUtils.GenerateImgHash(codeArchive)
	if err != nil {
		s.logger.Error("failed to generate code archive hash", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to generate code archive hash")
	}
	checkerImg := checkerImgDomain.CheckerImg{
		Uuid:        uuid.New(),
		Hash:        codeArchiveHash,
		CodeArchive: codeArchive,
	}
	if err := s.checkerService.CheckerImgService.Create(&checkerImg); err != nil {
		s.logger.Error("failed to create checker image", zap.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to create checker image")
	}
	response.Id = checkerImg.Uuid.String()

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

	vulnServiceList, err := s.vulnService.GetList(checkerUuid)
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
