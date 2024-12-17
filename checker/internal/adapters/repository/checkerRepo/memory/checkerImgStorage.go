package memory

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain/models"

	"sync"
)

type CheckerImgStorage struct {
	servicesData map[uuid.UUID]*models.VulnServiceImg
	mu           sync.Mutex
}

func NewCheckerImgStorage() *CheckerImgStorage {
	return &CheckerImgStorage{
		servicesData: make(map[uuid.UUID]*models.VulnServiceImg),
		mu:           sync.Mutex{},
	}
}

func (s *CheckerImgStorage) Add(codeArchive string) uuid.UUID {
	s.mu.Lock()
	defer s.mu.Unlock()
	id := uuid.New()
	s.servicesData[id] = &models.VulnServiceImg{
		Uuid:        id,
		CodeArchive: codeArchive,
	}
	return id
}
