package checkerRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
	models2 "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain/models"
)

type Repository interface {
	CreateChecker(checker *models.Checker) error
	UpdateChecker(checker *models.Checker) error
	GetChecker(UUID uuid.UUID) (*models.Checker, error)
	GetVulnServiceList(UUID uuid.UUID) ([]*models2.VulnService, error)
}
