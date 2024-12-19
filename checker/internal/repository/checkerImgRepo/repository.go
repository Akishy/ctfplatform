package checkerImgRepo

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
)

type Repository interface {
	CreateCheckerImg(checkerImg *checkerImgDomain.CheckerImg) error
	GetCheckerImg(checkerImgUuid uuid.UUID) (checkerImg *checkerImgDomain.CheckerImg, err error)
	CompareRawCheckerImg(raw string) (Img *checkerImgDomain.CheckerImg, err error)
}
