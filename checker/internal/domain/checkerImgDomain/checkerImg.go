package checkerImgDomain

import "github.com/google/uuid"

// образ чекера
type CheckerImg struct {
	Uuid        uuid.UUID
	Hash        string
	CodeArchive string
}
