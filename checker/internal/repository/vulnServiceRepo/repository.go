package vulnServiceRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain/models"
)

type Repository interface {
	GetVulnService(vulnServiceUUID uuid.UUID) (*models.VulnService, error)
	GetVulnServiceList(checkerUUID uuid.UUID) ([]*models.VulnService, error)
	CreateVulnService(vulnService *models.VulnService) error
	UpdateVulnService(vulnService *models.VulnService) error
}
