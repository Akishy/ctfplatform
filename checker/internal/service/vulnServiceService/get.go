package vulnServiceService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain/models"
)

func (s *Service) Get(UUID uuid.UUID) (*models.VulnService, error) {
	return s.repo.Get(UUID)
}
