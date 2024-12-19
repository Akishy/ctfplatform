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

func (s *Service) Check() error {
	checkers, err := s.ListAllRegistered()
	if err != nil {
		return err
	}
	client := http.Client{
		Timeout: time.Second * 3, // задать глобальную переменную defaultTimeout
	}
	wg := &sync.WaitGroup{}

	for _, checker := range checkers {
		go func() {
			wg.Add(1)
			defer wg.Done()
			vulnServices, err := s.repo.GetVulnServiceList(checker.UUID)
			if err != nil {
				s.logger.Warn("No vulnServices")
				//s.logger.Error("GetVulnServiceList", zap.Error(err))
			}
			innerwg := &sync.WaitGroup{}
			for _, vulnService := range vulnServices {
				go func() {
					innerwg.Add(1)
					defer innerwg.Done()
					reqUuid := uuid.New()
					s.logger.Debug("reqUUID generated", zap.Any("uuid", reqUuid.String()))
					flag := s.FlagGeneratorService.Generate()
					s.logger.Debug("flag generated", zap.Any("flag", flag))

					request := checkRequest{
						RequestUUID:     reqUuid,
						VulnServiceIP:   vulnService.Ip,
						VulnServicePort: vulnService.WebPort,
						Flag:            flag.Flag.String(),
					}
					s.logger.Debug("request struct created", zap.Any("request", request))

					bytesBody, err := json.Marshal(request)
					if err != nil {
						s.logger.Error("(checkService) failed to marshal check request", zap.Error(err))
						return
					}
					s.logger.Debug("check request marshalled", zap.Any("request", request))
					reqBody := bytes.NewReader(bytesBody)

					resp, err := client.Post(fmt.Sprintf("http://%v:%v/checkVulnService", checker.Ip, checker.WebPort), "application/json", reqBody)
					if err != nil {
						s.logger.Error("(checkService) failed to send check request", zap.Error(err))
						return
					}

					var response checkResponse
					err = json.NewDecoder(resp.Body).Decode(&response)
					if err != nil {
						s.logger.Error("(checkService) failed to unmarshal check response", zap.Error(err))
						return
					}

					if !response.IsTaskAccepted {
						s.logger.Error("(checkService) task is not accepted")
						return
					}

					if s.repo.CreateRequestToVulnService(request.RequestUUID, vulnService.Uuid) != nil {
						s.logger.Error("(checkService) failed to connect request id with vuln service ", zap.Error(err))
					}

					if s.repo.CreateFlag(flag) != nil {
						s.logger.Error("(checkService) failed to create flag", zap.Error(err))
					}
				}()
			}
			innerwg.Wait()
		}()
		wg.Wait()
	}

	return nil
}
