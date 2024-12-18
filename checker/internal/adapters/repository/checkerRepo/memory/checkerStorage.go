package memory

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
	models2 "gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/vulnServiceDomain/models"
	"io"
	"net/http"
	"sync"
)

type CheckerStorage struct {
	checkersData map[uuid.UUID]*models.Checker
	mu           sync.Mutex
}

func NewCheckerStorage() *CheckerStorage {
	return &CheckerStorage{
		checkersData: make(map[uuid.UUID]*models.Checker),
		mu:           sync.Mutex{},
	}
}

func (storage *CheckerStorage) CreateChecker(checker *models.Checker) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	id := uuid.New()
	storage.checkersData[id] = checker
	return nil
}

func (storage *CheckerStorage) UpdateChecker(checker *models.Checker) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	// unimplemented
	return errors.New("not yet implemented")
}

func (storage *CheckerStorage) GetChecker(id uuid.UUID) (*models.Checker, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	checker, ok := storage.checkersData[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("checker [%s] not found", id.String()))
	}
	return checker, nil
}

func (storage *CheckerStorage) GetVulnServiceList(checkerUUID uuid.UUID) ([]*models2.VulnService, error) {

	return nil, errors.New(fmt.Sprintf("Cannot get vuln service status, %v", err))

}
