package memory

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	models2 "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain/models"
	"sync"
)

type Storage struct {
	vulnServicesData map[uuid.UUID]*models.VulnService
	checkersData     map[uuid.UUID]*models2.Checker
	mu               sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		vulnServicesData: make(map[uuid.UUID]*models.VulnService),
		checkersData:     make(map[uuid.UUID]*models2.Checker),
		mu:               sync.RWMutex{},
	}
}

func (s *Storage) GetVulnService(vulnServiceUUID uuid.UUID) (*models.VulnService, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	vulnService, ok := s.vulnServicesData[vulnServiceUUID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("vulnService [%v] not found", vulnServiceUUID.String()))
	}
	return vulnService, nil
}

func (s *Storage) GetVulnServiceList(checkerUUID uuid.UUID) ([]*models.VulnService, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	vulnServiceList := make([]*models.VulnService, 0)
	for _, vulnService := range s.vulnServicesData {
		if vulnService.CheckerId == checkerUUID {
			vulnServiceList = append(vulnServiceList, vulnService)
		}
	}
	return vulnServiceList, nil
}

func (s *Storage) CreateVulnService(vulnService *models.VulnService) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.vulnServicesData[vulnService.Uuid]; ok {
		return errors.New(fmt.Sprintf("vulnService [%v] already exists", vulnService.Uuid.String()))
	}
	s.vulnServicesData[vulnService.Uuid] = vulnService
	return nil
}

func (s *Storage) UpdateVulnService(vulnService *models.VulnService) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.vulnServicesData[vulnService.Uuid]; !ok {
		return errors.New(fmt.Sprintf("vulnService [%v] not found", vulnService.Uuid.String()))
	}

	s.vulnServicesData[vulnService.Uuid] = vulnService
	return nil
}
