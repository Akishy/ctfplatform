package checkerService

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/zap"
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
				s.logger.Error("(checkService) failed to marshal check request", zap.Error(err))
			}
			reqBody := bytes.NewReader(bytesBody)

			resp, err := client.Post(fmt.Sprintf("http://%v:%v/checkVulnService", checker.Ip, checker.WebPort), "application/json", reqBody)
			if err != nil {
				s.logger.Error("(checkService) failed to send check request", zap.Error(err))
			}

			var response checkResponse
			err = json.NewDecoder(resp.Body).Decode(&response)
			if err != nil {
				s.logger.Error("(checkService) failed to unmarshal check response", zap.Error(err))
			}
			if !response.IsTaskAccepted {
				s.logger.Error("(checkService) task is not accepted")
			}

			if s.repo.CreateRequestToVulnService(request.RequestUUID, vulnService.Uuid) != nil {
				s.logger.Error("(checkService) failed to connect request id with vuln service ", zap.Error(err))
			}

			if s.repo.CreateFlag(flag) != nil {
				s.logger.Error("(checkService) failed to create flag", zap.Error(err))
			}
		}()

	}
	defer wg.Wait()

	return nil
}
