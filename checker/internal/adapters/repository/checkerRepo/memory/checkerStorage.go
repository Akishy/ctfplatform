package memory

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
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
