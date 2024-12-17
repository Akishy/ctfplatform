package models

import "github.com/google/uuid"

// образ уязвимого сервиса
type VulnServiceImg struct {
	Uuid        uuid.UUID
	CodeArchive string
}
