package checkerService

import (
	"bytes"
	"encoding/json"
	"errors"
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
		Timeout: time.Second * 3, // задать глобальную переменную defaultTimeout
	}
	wg := &sync.WaitGroup{}

	errCh := make(chan error) // implement errgroup

	for _, vulnService := range vulnServices {
		// а потом чекер будет отправлять результат по ручке /checker/sendVulnServiceStatus
		go func() {
			wg.Add(1)
			defer wg.Done()
			reqUuid := uuid.New()
			flag := s.flagGenerator.Generate()

			request := checkRequest{
				RequestUUID:     reqUuid,
				VulnServiceIP:   vulnService.Ip,
				VulnServicePort: vulnService.WebPort,
				Flag:            flag.Flag.String(),
			}

			bytesBody, err := json.Marshal(request)
			if err != nil {
				errCh <- err
				return
			}
			reqBody := bytes.NewReader(bytesBody)

			resp, err := client.Post(fmt.Sprintf("http://%v:%v/checkVulnService", checker.Ip, checker.WebPort), "application/json", reqBody)
			if err != nil {
				errCh <- err
				return
			}

			var response checkResponse
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				errCh <- err
				return
			}
			if !response.IsTaskAccepted {
				errCh <- errors.New("task was not accepted by checker")
			}

			if s.repo.CreateRequestToVulnService(request.RequestUUID, vulnService.Uuid) != nil {
				errCh <- err
				return
			}

			if s.repo.CreateFlag(flag) != nil {
				errCh <- err
				return
			}
		}()

	}
	defer wg.Wait()
	// select
	fnError := <-errCh
	return fnError
}
