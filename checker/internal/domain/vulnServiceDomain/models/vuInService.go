package models

import "github.com/google/uuid"

// конкретный сервис участиков
type VulnService struct {
	Uuid          uuid.UUID
	Ip            string
	WebPort       int
	StatusCode    int
	Message       string
	VulnServiceId uuid.UUID
}
