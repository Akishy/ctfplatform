package checkerRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
)

type Repository interface {
	Create(checker *models.Checker) error
	Update(checker *models.Checker) error
	GetVulnServiceStatus(UUID uuid.UUID) (string, error)
}
