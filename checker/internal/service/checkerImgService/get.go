package checkerImgService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"

func (s *Service) Get(imgHash string) (*checkerImgDomain.CheckerImg, error) {
	return s.repo.GetCheckerImg(imgHash)
}
