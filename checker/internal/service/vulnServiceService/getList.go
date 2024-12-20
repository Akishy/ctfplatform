package vulnServiceService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

// get list of services by checker uuid
func (s *Service) GetActiveList(checkerUUID uuid.UUID) ([]*vulnServiceDomain.VulnService, error) {
	return s.repo.GetActiveVulnServiceList(checkerUUID)
}
