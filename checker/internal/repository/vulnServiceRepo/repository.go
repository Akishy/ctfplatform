package vulnServiceRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

type Repository interface {
	GetVulnService(vulnServiceUUID uuid.UUID) (*vulnServiceDomain.VulnService, error)
	GetVulnServiceList(checkerUUID uuid.UUID) ([]*vulnServiceDomain.VulnService, error)
	CreateVulnService(vulnService *vulnServiceDomain.VulnService) error
	UpdateVulnService(vulnService *vulnServiceDomain.VulnService) error
}
