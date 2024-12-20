package vulnServiceService

import (
	"github.com/google/uuid"
)

func (s *Service) DeactivateVulnService(UUID uuid.UUID) error {
	_, err := s.Get(UUID)
	if err != nil {
		return err
	}

	return s.repo.DeactivateVulnService(UUID)
}
