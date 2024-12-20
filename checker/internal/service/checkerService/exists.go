package checkerService

import "github.com/google/uuid"

func (s *Service) Exists(checkerUUID uuid.UUID) bool {
	if _, err := s.repo.GetChecker(checkerUUID); err != nil {
		return false
	}
	return true
}
