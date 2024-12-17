package memory

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vuInServiceDomain/models"

	"sync"
)

type VuInServiceStorage struct {
	servicesData map[uuid.UUID]*models.VuInService
	mu           sync.Mutex
}

func NewVuInServiceStorage() *VuInServiceStorage {
	return &VuInServiceStorage{
		servicesData: make(map[uuid.UUID]*models.VuInService),
		mu:           sync.Mutex{},
	}
}

func (s *VuInServiceStorage) Add(codeArchive string) uuid.UUID {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New()
	s.servicesData[id] = &models.VuInService{
		Uuid:        id,
		CodeArchive: codeArchive,
	}
	return id
}
