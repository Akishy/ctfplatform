package checkerService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"

func (s *Service) Create(checker *models.Checker) error {
	return s.repo.CreateChecker(checker)
}
