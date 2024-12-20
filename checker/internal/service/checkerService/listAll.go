package checkerService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"

func (s *Service) ListAllRegistered() ([]*checkerDomain.Checker, error) {
	checkers, err := s.repo.ListAllRegisteredCheckers()
	if err != nil {
		return nil, err
	}
	return checkers, nil
}
