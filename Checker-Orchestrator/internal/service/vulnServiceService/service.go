package vulnServiceService

import vulnServiceRepo "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/repository/vulnServiceRepo"

//

type Service struct {
	repo vulnServiceRepo.Repository
}

func NewService(repository vulnServiceRepo.Repository) *Service {
	return &Service{
		repo: repository,
	}
}
