package memory

import (
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
	"sync"
)

type Storage struct {
	checkersData map[uuid.UUID]*models.Checker
	servicesData map[uuid.UUID]interface{} // edit!!!
	mu           sync.Mutex
}

func New() *Storage {
	return &Storage{
		checkersData: make(map[uuid.UUID]*models.Checker),
		servicesData: make(map[uuid.UUID]interface{}),
		mu:           sync.Mutex{},
	}
}

func (storage *Storage) CreateChecker(checker *models.Checker) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	id := uuid.New()
	storage.checkersData[id] = checker
	return nil
}
