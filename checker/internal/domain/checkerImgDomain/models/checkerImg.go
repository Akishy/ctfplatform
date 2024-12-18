package models

import "github.com/google/uuid"

// образ чекера
type CheckerImg struct {
	Uuid        uuid.UUID
	CodeArchive string
}
