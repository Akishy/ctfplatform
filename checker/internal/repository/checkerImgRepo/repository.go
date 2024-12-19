package checkerImgRepo

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"

type Repository interface {
	CreateCheckerImg(checkerImg *checkerImgDomain.CheckerImg) error
	GetCheckerImg(ImgHash string) (checkerImg *checkerImgDomain.CheckerImg, err error)
}
