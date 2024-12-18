package models

import "github.com/google/uuid"

// конкретный сервис участиков
type VulnService struct {
	Uuid       uuid.UUID // uuid сервиса
	Ip         string
	WebPort    int
	StatusCode int
	Message    string
	CheckerId  uuid.UUID
}
