package memory

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vuInServiceInstanceDomain/models"

	"sync"
)

type VuInServiceInstanceStorage struct {
	servicesData map[uuid.UUID]*models.VuInServiceInstance
	mu           sync.Mutex
}

func NewVuInServiceInstanceStorage() *VuInServiceInstanceStorage {
	return &VuInServiceInstanceStorage{
		servicesData: make(map[uuid.UUID]*models.VuInServiceInstance),
		mu:           sync.Mutex{},
	}
}

func (s *VuInServiceInstanceStorage) Add(serviceInstance *models.VuInServiceInstance) {

}
