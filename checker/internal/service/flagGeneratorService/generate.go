package flagGeneratorService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain/models"
)

func (s *Service) Generate() models.Flag {
	flag := models.Flag{}
	uuid := uuid.New()
	flag.Flag = uuid
	flag.Status = models.FLAG_PUSHED_TO_CHECKER

}
