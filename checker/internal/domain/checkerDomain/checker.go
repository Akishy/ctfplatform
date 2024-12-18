package checkerDomain

import (
	"github.com/google/uuid"
)

type Checker struct {
	UUID uuid.UUID
	//CheckerImg
	Ip      string
	WebPort int
}
