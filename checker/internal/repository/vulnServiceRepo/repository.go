package vulnServiceRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
)

type Repository interface {
	vulnServiceRepo
	vulnServiceRequestRepo
}

type vulnServiceRepo interface {
	GetVulnService(vulnServiceUUID uuid.UUID) (*vulnServiceDomain.VulnService, error)
	GetActiveVulnServiceList(checkerUUID uuid.UUID) ([]*vulnServiceDomain.VulnService, error)
	CreateVulnService(vulnService *vulnServiceDomain.VulnService) error
	UpdateVulnService(uuid uuid.UUID, statusCode vulnServiceDomain.VulnServiceStatus, message string, lastCheck int64) error
	DeactivateVulnService(uuid uuid.UUID) error
}
type vulnServiceRequestRepo interface {
	GetRequestToVulnService(requestUUID uuid.UUID) (*vulnServiceDomain.RequestToVulnService, error)
	CreateRequestToVulnService(requestUUID, vulnServiceUUID uuid.UUID) error
}
