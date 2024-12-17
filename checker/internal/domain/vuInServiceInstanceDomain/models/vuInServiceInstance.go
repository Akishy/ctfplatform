package models

import "github.com/google/uuid"

type VuInServiceInstance struct {
	Uuid          uuid.UUID
	Ip            string
	WebPort       int
	StatusCode    int
	Message       string
	VulnServiceId uuid.UUID
}
