package vulnServiceService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

func (s *Service) UpdateByRequestUUID(UUID uuid.UUID,
	statusCode vulnServiceDomain.VulnServiceStatus,
	message string,
	lastCheck int64) error {

	vulnServiceId, err := s.repo.GetRequestToVulnService(UUID)
	if err != nil {
		return err
	}

	vulnService, err := s.Get(vulnServiceId.VulnServiceId)
	if err != nil {
		return err
	}

	return s.repo.UpdateVulnService(vulnService.Uuid, statusCode, message, lastCheck)
}
