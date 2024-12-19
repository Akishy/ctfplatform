package checkerDomain

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
)

type Checker struct {
	UUID       uuid.UUID
	CheckerImg *checkerImgDomain.CheckerImg
	Ip         string
	WebPort    int
}
