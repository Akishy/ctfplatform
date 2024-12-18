package models

import "github.com/google/uuid"

type VulnServiceStatus uint

const (
	OK VulnServiceStatus = iota // доставлен в чекер
	MUMBLE
	CORRUPT
	DOWN
)

// VulnService - конкретный сервис участиков
type VulnService struct {
	Uuid       uuid.UUID // uuid сервиса
	Ip         string
	WebPort    int
	StatusCode VulnServiceStatus
	Message    string
	CheckerId  uuid.UUID
}
