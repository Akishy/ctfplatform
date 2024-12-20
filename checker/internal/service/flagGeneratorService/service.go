package flagGeneratorService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/flagGeneratorRepo"

type Service struct {
	Repo flagGeneratorRepo.Repository
}

func NewService(flagGeneratorRepo flagGeneratorRepo.Repository) *Service {
	return &Service{
		Repo: flagGeneratorRepo,
	}
}
