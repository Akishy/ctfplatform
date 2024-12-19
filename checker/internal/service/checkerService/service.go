package checkerService

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/checkerRepo"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/flagGeneratorService"
	"go.uber.org/zap"
)

type Service struct {
	repo          checkerRepo.Repository
	flagGenerator *flagGeneratorService.Service
	logger        *zap.Logger
}

func NewService(checkerRepo checkerRepo.Repository, flagGeneratorService *flagGeneratorService.Service, logger *zap.Logger) *Service {
	return &Service{
		repo:          checkerRepo,
		flagGenerator: flagGeneratorService,
		logger:        logger,
	}
}
