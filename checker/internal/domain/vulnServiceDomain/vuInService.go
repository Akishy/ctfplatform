package vulnServiceDomain

import (
	"github.com/google/uuid"
)

type VulnServiceStatus uint

const (
	OK VulnServiceStatus = iota // Доставлен в чекер
	MUMBLE
	CORRUPT
	DOWN
)

// VulnService - конкретный сервис участников
type VulnService struct {
	Uuid       uuid.UUID // uuid сервиса
	Ip         string
	WebPort    int
	StatusCode VulnServiceStatus
	Message    string
	CheckerId  uuid.UUID
	LastCheck  int64
}
