package vulnServiceService

import "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"

func (s *Service) Create(vulnService *vulnServiceDomain.VulnService) error {
	return s.repo.CreateVulnService(vulnService)
}
