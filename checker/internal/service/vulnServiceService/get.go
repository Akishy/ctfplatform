package vulnServiceService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

func (s *Service) Get(UUID uuid.UUID) (*vulnServiceDomain.VulnService, error) {
	return s.repo.Get(UUID)
}
