package flagGeneratorRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain"
)

type Repository interface {
	CreateFlag(flag *flagGeneratorDomain.Flag) error
	GetFlagInfo(uuid uuid.UUID) (flagGeneratorDomain.Flag, error)
}
