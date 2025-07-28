package checkerImgService

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/checkerImgRepo"
	"go.uber.org/zap"
)

type Service struct {
	repo   checkerImgRepo.Repository
	logger *zap.Logger
}

func NewService(checkerImgRepo checkerImgRepo.Repository, logger *zap.Logger) *Service {
	return &Service{
		repo:   checkerImgRepo,
		logger: logger,
	}
}
