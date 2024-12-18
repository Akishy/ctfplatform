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
	vulnServicesData           map[uuid.UUID]*models.VulnService
	checkersData               map[uuid.UUID]*models2.Checker
	requestsToVulnServicesData map[uuid.UUID]*models.RequestToVulnService
	mu                         sync.RWMutex
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
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.vulnServicesData[vulnService.Uuid]; ok {
		return errors.New(fmt.Sprintf("vulnService [%v] already exists", vulnService.Uuid.String()))
	}
	s.vulnServicesData[vulnService.Uuid] = vulnService
	return nil
}

func (s *Storage) UpdateVulnService(vulnService *models.VulnService) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.vulnServicesData[vulnService.Uuid]; !ok {
		return errors.New(fmt.Sprintf("vulnService [%v] not found", vulnService.Uuid.String()))
	}

	s.vulnServicesData[vulnService.Uuid] = vulnService
	return nil
}

// хз в какой интерфейс, скорее всего VulnService
func (s *Storage) CreateRequestToVulnService(requestUUID, vulnServiceUUID uuid.UUID) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.requestsToVulnServicesData[requestUUID]; ok {
		return errors.New(fmt.Sprintf("vulnService [%v] already exists", vulnServiceUUID.String()))
	}
	s.requestsToVulnServicesData[requestUUID] = &models.RequestToVulnService{
		Uuid:          requestUUID,
		VulnServiceId: vulnServiceUUID,
	}
	return nil
}

// not implemented
//func (s *Storage) CreateChecker(checker *models2.Checker) error        { return nil }
//func (s *Storage) UpdateChecker(checker *models2.Checker) error        { return nil }
//func (s *Storage) GetChecker(UUID uuid.UUID) (*models2.Checker, error) { return nil, nil }
