package checkerImgService

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
)

func (s *Service) GetByRaw(rawImg string) (Img *checkerImgDomain.CheckerImg, err error) {
	return s.repo.CompareRawCheckerImg(rawImg)
}
