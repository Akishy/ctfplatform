package checkerService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"sync"
	"time"
)

func (s *Service) Check(checkerUUID uuid.UUID) error {
	checker, err := s.repo.GetChecker(checkerUUID)
	if err != nil {
		return err
	}

	vulnServices, err := s.repo.GetVulnServiceList(checker.UUID)
	if err != nil {
		return err
	}

	client := http.Client{
		Timeout: time.Second * 3,
	}
	wg := &sync.WaitGroup{}

	for _, vulnService := range vulnServices {
		// а потом чекер будет отправлять результат по ручке /checker/sendVulnServiceStatus
		go func() {
			wg.Add(1)
			defer wg.Done()
			reqUuid := uuid.New()

			request := checkRequest{
				RequestUUID:     reqUuid,
				VulnServiceIP:   vulnService.Ip,
				VulnServicePort: vulnService.WebPort,
				Flag:            "testFlag",
			}

			bytesBody, err := json.Marshal(request)
			if err != nil {
				//
				// Что делать при ошибке?
				//
			}
			reqBody := bytes.NewReader(bytesBody)

			resp, err := client.Post(fmt.Sprintf("http://%v:%v/checkVulnService", checker.Ip, checker.WebPort), "application/json", reqBody)
			if err != nil {
				//
				// Что делать при ошибке?
				//
			}

			var response checkResponse
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				//
				// Что делать при ошибке?
				//
			}

		}()

	}
	defer wg.Wait()

	//
	//...
	//

	return nil
}
