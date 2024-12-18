package checkerService

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/checkerRepo"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/flagGeneratorService"
)

type Service struct {
	repo          checkerRepo.Repository
	flagGenerator *flagGeneratorService.Service
}

func NewService(checkerRepo checkerRepo.Repository, flagGeneratorService *flagGeneratorService.Service) *Service {
	return &Service{
		repo:          checkerRepo,
		flagGenerator: flagGeneratorService,
	}
}
