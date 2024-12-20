package vulnServiceDomain

import "github.com/google/uuid"

type RequestToVulnService struct {
	Uuid          uuid.UUID
	VulnServiceId uuid.UUID
}
