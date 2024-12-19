package checkerImgService

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
)

func (s *Service) Get(imgUuid uuid.UUID) (*checkerImgDomain.CheckerImg, error) {
	return s.repo.GetCheckerImg(imgUuid)
}
