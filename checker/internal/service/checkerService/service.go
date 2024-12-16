package checkerService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/checkerRepo"

type Service struct {
	repo checkerRepo.Repository
}

func NewService(checkerRepo checkerRepo.Repository) *Service {
	return &Service{
		repo: checkerRepo,
	}
}
