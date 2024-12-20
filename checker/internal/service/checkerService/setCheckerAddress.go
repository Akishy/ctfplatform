package checkerService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
	"go.uber.org/zap"
)

func (s *Service) SetCheckerAddress(uuid uuid.UUID, ip string, port int) error {
	checker := checkerDomain.Checker{
		CheckerImg: nil, // найти и присвоить
		Ip:         "",
		WebPort:    0,
	}
	if checkerPointer, err := s.repo.GetChecker(uuid); err != nil {
		s.logger.Error("failed to get checker by uuid", zap.Error(err))
		return err
	} else {
		checker = *checkerPointer // копирование чтоб не изменять структуру тут (работа репо)
	}
	checker.Ip = ip
	checker.WebPort = port

	return s.repo.UpdateChecker(&checker)
}
