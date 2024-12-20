package checkerService

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/checkerRepo"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/checkerImgService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/flagGeneratorService"
	"go.uber.org/zap"
)

type Service struct {
	repo                 checkerRepo.Repository
	FlagGeneratorService *flagGeneratorService.Service
	CheckerImgService    *checkerImgService.Service
	logger               *zap.Logger
}

func NewService(checkerRepo checkerRepo.Repository, flagGeneratorService *flagGeneratorService.Service, checkerImgService *checkerImgService.Service, logger *zap.Logger) *Service {
	return &Service{
		repo:                 checkerRepo,
		FlagGeneratorService: flagGeneratorService,
		CheckerImgService:    checkerImgService,
		logger:               logger,
	}
}
