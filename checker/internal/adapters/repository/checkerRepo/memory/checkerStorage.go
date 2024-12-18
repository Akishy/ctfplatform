package memory

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/domain/checkerDomain/models"
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

func (storage *CheckerStorage) Create(checker *models.Checker) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	id := uuid.New()
	storage.checkersData[id] = checker
	return nil
}

func (storage *CheckerStorage) Update(checker *models.Checker) error {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	// unimplemented
	return errors.New("not yet implemented")
}

func (storage *CheckerStorage) Get(id uuid.UUID) (*models.Checker, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	checker, ok := storage.checkersData[id]
	if !ok {
		return nil, errors.New(fmt.Sprintf("checker [%s] not found", id.String()))
	}
	return checker, nil
}

func (storage *CheckerStorage) GetVulnServiceStatus(id uuid.UUID) (string, error) {
	storage.mu.Lock()
	defer storage.mu.Unlock()
	checker, err := storage.Get(id)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Cannot get vuln service status, %v", err))
	}
	// пойти к чекеру с запросом
	// for i in range (количество уязвимых сервисов, которые может проверить этот чекер)
	// я не понимаю как понять, может ли чекер проверить конкретный уязвимый сервис
	client := http.Client{}
	uuid := uuid.New()
	data := byte[fmt.Sprintf(`{"request_UUID":"%v"}`, uuid.String())]
	var requestBody io.Reader

	resp, err := client.Post(fmt.Sprintf("http://%v:%v/checkVulnService", checker.Ip, checker.WebPort), "application/json", nil)
	// request UUID
	//IP VulnService
	// Port VulnService
	// Flag

}
