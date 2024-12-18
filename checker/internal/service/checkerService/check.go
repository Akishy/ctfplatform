package checkerService

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"sync"
)

func (s *Service) Check(checkerUUID uuid.UUID, flag string) (string, error) {
	checker, err := s.repo.GetChecker(checkerUUID)
	if err != nil {
		return "", err
	}

	vulnServices, err := s.repo.GetVulnServiceList(checker.UUID)
	if err != nil {
		return "", err
	}

	// sync.WaitGroup?
	client := http.Client{}
	wg := &sync.WaitGroup{}

	for _, vulnService := range vulnServices {
		// получается тут делаем http client
		// в горутинах?
		// ответом на эту хуйню будет статус: принято, так что ждём
		// а потом чекер будет отправлять результат по ручке /checker/sendVulnServiceStatus
		go func() {
			wg.Add(1)
			defer wg.Done()
			reqUuid := uuid.New()
			// request UUID
			//IP VulnService
			// Port VulnService
			// Flag
			// нет, надо не один флаг. Каждому сервису уникальный. Тогда генератор передаем. Я бы в структуру прям записал его. В структуру Checker? CheckerSystem то есть, да?
			data := []byte(fmt.Sprintf(`{"request_UUID":"%v","vuln_service_ip":"%v","vuln_service_port":"%v","flag":"%v"}`, reqUuid.String(), vulnService.Ip, vulnService.WebPort, "testFlag"))
			reqBody := bytes.NewReader(data)

			resp, err := client.Post(fmt.Sprintf("http://%v:%v/checkVulnService", checker.Ip, checker.WebPort), "application/json", reqBody)

		}()

	}
	defer wg.Wait()
}
