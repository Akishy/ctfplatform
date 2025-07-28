package checkerService

import "github.com/google/uuid"

func (s *Service) Subscribed(checkerUUID uuid.UUID) bool {
	if checker, err := s.repo.GetChecker(checkerUUID); err != nil {
		return false
	} else {
		if checker.Ip == "" && checker.WebPort == 0 {
			return false
		}
	}
	return true
}
