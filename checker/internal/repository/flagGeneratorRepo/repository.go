package flagGeneratorRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain/models"
)

type Repository interface {
	CreateFlag(flag *models.Flag) error
	GetFlagInfo(uuid uuid.UUID) (models.Flag, error)
}
