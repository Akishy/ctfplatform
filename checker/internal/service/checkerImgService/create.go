package checkerImgService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"

func (s *Service) Create(checkerImg *checkerImgDomain.CheckerImg) error {
	return s.repo.CreateCheckerImg(checkerImg)
}
