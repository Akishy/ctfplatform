package flagGeneratorService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain"
)

func (s *Service) Generate() *flagGeneratorDomain.Flag {
	return &flagGeneratorDomain.Flag{
		UUID:   uuid.New(),
		Flag:   uuid.New(),
		Status: flagGeneratorDomain.FLAG_PUSHED_TO_CHECKER,
	}
}
