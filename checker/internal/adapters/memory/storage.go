package memory

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerImgDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/flagGeneratorDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/hashUtils"
	"sync"
)

type Storage struct {
	vulnServicesData           map[uuid.UUID]*vulnServiceDomain.VulnService
	checkersData               map[uuid.UUID]*checkerDomain.Checker
	requestsToVulnServicesData map[uuid.UUID]*vulnServiceDomain.RequestToVulnService
	checkerImgData             map[uuid.UUID]*checkerImgDomain.CheckerImg
	flagData                   map[uuid.UUID]*flagGeneratorDomain.Flag
	mu                         sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		vulnServicesData:           make(map[uuid.UUID]*vulnServiceDomain.VulnService),
		checkersData:               make(map[uuid.UUID]*checkerDomain.Checker),
		requestsToVulnServicesData: make(map[uuid.UUID]*vulnServiceDomain.RequestToVulnService),
		checkerImgData:             make(map[uuid.UUID]*checkerImgDomain.CheckerImg),
		flagData:                   make(map[uuid.UUID]*flagGeneratorDomain.Flag),
		mu:                         sync.RWMutex{},
	}
}

func (s *Storage) GetVulnService(vulnServiceUUID uuid.UUID) (*vulnServiceDomain.VulnService, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	vulnService, ok := s.vulnServicesData[vulnServiceUUID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("vulnService [%v] not found", vulnServiceUUID.String()))
	}
	return vulnService, nil
}

func (s *Storage) GetVulnServiceList(checkerUUID uuid.UUID) ([]*vulnServiceDomain.VulnService, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	vulnServiceList := make([]*vulnServiceDomain.VulnService, 0)
	for _, vulnService := range s.vulnServicesData {
		if vulnService.CheckerId == checkerUUID {
			vulnServiceList = append(vulnServiceList, vulnService)
		}
	}
	if len(vulnServiceList) == 0 {
		return nil, errors.New(fmt.Sprintf("vulnService list is empty"))
	}
	return vulnServiceList, nil
}

func (s *Storage) CreateVulnService(vulnService *vulnServiceDomain.VulnService) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.vulnServicesData[vulnService.Uuid]; ok {
		return errors.New(fmt.Sprintf("vulnService [%v] already exists", vulnService.Uuid.String()))
	}
	s.vulnServicesData[vulnService.Uuid] = vulnService
	return nil
}

func (s *Storage) UpdateVulnService(
	uuid uuid.UUID,
	statusCode vulnServiceDomain.VulnServiceStatus,
	message string,
	lastCheck int64) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if vulnService, ok := s.vulnServicesData[uuid]; !ok {
		return errors.New(fmt.Sprintf("vulnService [%v] not found", vulnService.Uuid.String()))
	} else {
		vulnService.StatusCode = statusCode
		vulnService.Message = message
		vulnService.LastCheck = lastCheck
	}

	return nil
}

func (s *Storage) CreateChecker(checker *checkerDomain.Checker) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.checkersData[checker.UUID]; ok {
		return errors.New(fmt.Sprintf("checker [%v] already exists", checker.UUID.String()))
	}
	s.checkersData[checker.UUID] = checker
	return nil
}
func (s *Storage) UpdateChecker(checker *checkerDomain.Checker) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.checkersData[checker.UUID]; !ok {
		return errors.New(fmt.Sprintf("checker [%v] not found", checker.UUID.String()))
	}
	s.checkersData[checker.UUID] = checker
	return nil
}
func (s *Storage) GetChecker(UUID uuid.UUID) (*checkerDomain.Checker, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	checker, ok := s.checkersData[UUID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("checker [%v] not found", UUID.String()))
	}
	return checker, nil
}

// хз в какой интерфейс, скорее всего VulnService
func (s *Storage) CreateRequestToVulnService(requestUUID, vulnServiceUUID uuid.UUID) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if _, ok := s.requestsToVulnServicesData[requestUUID]; ok {
		return errors.New(fmt.Sprintf("vulnService [%v] already exists", vulnServiceUUID.String()))
	}
	s.requestsToVulnServicesData[requestUUID] = &vulnServiceDomain.RequestToVulnService{
		Uuid:          requestUUID,
		VulnServiceId: vulnServiceUUID,
	}
	return nil
}

func (s *Storage) GetRequestToVulnService(requestUUID uuid.UUID) (*vulnServiceDomain.RequestToVulnService, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	requestToVulnService, ok := s.requestsToVulnServicesData[requestUUID]
	if !ok {
		return nil, errors.New(fmt.Sprintf("requestToVulnService [%v] not found", requestUUID.String()))
	}
	return requestToVulnService, nil
}

func (s *Storage) CreateFlag(flag *flagGeneratorDomain.Flag) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.flagData[flag.UUID]; ok {
		return errors.New(fmt.Sprintf("flag [%v] already exists", flag.UUID.String()))
	}
	s.flagData[flag.UUID] = flag
	return nil
}

func (s *Storage) GetFlagInfo(uuid uuid.UUID) (*flagGeneratorDomain.Flag, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	flagInfo, ok := s.flagData[uuid]
	if !ok {
		return nil, errors.New(fmt.Sprintf("flag [%v] not found", uuid.String()))
	}
	return flagInfo, nil
}

func (s *Storage) CreateCheckerImg(checkerImg *checkerImgDomain.CheckerImg) error {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if _, ok := s.checkerImgData[checkerImg.Uuid]; ok {
		return errors.New(fmt.Sprintf("checkerImg [%v] already exists", checkerImg.Hash))
	}
	s.checkerImgData[checkerImg.Uuid] = checkerImg
	return nil
}

func (s *Storage) GetCheckerImg(checkerImgUuid uuid.UUID) (checkerImg *checkerImgDomain.CheckerImg, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	checkerImg, ok := s.checkerImgData[checkerImgUuid]
	if !ok {
		return nil, errors.New(fmt.Sprintf("checkerImg [%v] not found", checkerImgUuid))
	}
	return checkerImg, nil
}

func (s *Storage) ListAllRegisteredCheckers() ([]*checkerDomain.Checker, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	checkers := make([]*checkerDomain.Checker, 0, len(s.checkersData))
	for _, checker := range s.checkersData {
		if checker.Ip != "" && checker.WebPort != 0 {
			checkers = append(checkers, checker)
		}
	}
	return checkers, nil
}

func (s *Storage) CompareRawCheckerImg(raw string) (Img *checkerImgDomain.CheckerImg, err error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, checkerImg := range s.checkerImgData {
		found := hashUtils.CompareImgHash(raw, checkerImg.Hash)
		if found {
			return checkerImg, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("checkerImg with byteslen [%v] not found", len(raw)))
}

//func (s *Storage) GetFlag(uuid uuid.UUID) (*flagGeneratorDomain.Flag, error) {
//
//}
