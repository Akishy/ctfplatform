package memory

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain/models"

	"sync"
)

type VulnServiceStorage struct {
	servicesData map[uuid.UUID]*models.VulnService
	mu           sync.Mutex
}

func NewVulnServiceStorage() *VulnServiceStorage {
	return &VulnServiceStorage{
		servicesData: make(map[uuid.UUID]*models.VulnService),
		mu:           sync.Mutex{},
	}
}

func (s *VulnServiceStorage) Add(serviceInstance *models.VulnService) {

}

func (s *VulnServiceStorage) GetList(checkerID uuid.UUID) ([]*models.VulnService, error) {
	
}
