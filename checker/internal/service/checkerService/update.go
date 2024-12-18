package checkerService

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
)

func (s *Service) Update(checker *checkerDomain.Checker) error {
	return s.repo.UpdateChecker(checker)
}
