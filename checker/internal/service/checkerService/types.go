package checkerService

import "github.com/google/uuid"

type checkRequest struct {
	RequestUUID     uuid.UUID `json:"request_uuid"`
	VulnServiceIP   string    `json:"vuln_service_ip"`
	VulnServicePort int       `json:"vuln_service_port"`
	Flag            string    `json:"flag"`
}

type checkResponse struct {
	IsTaskAccepted bool `json:"is_task_accepted"`
}
