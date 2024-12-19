package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/jsonUtils"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/checkerService"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/service/vulnServiceService"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ServiceServer struct {
	server             *http.Server
	logger             *zap.Logger
	VulnServiceService *vulnServiceService.Service
	CheckerService     *checkerService.Service
}

// New создает новый экземпляр ServiceServer
func New(port int, vulnServiceService *vulnServiceService.Service, logger *zap.Logger) *ServiceServer {
	router := mux.NewRouter()
	servServer := &ServiceServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		logger:             logger,
		VulnServiceService: vulnServiceService,
	}
	checkerRouter := router.PathPrefix("/checker").Subrouter()

	// Регистрация маршрута для subscribe
	checkerRouter.HandleFunc("/subscribe", servServer.SubscribeHandler).Methods("POST")

	// Регистрация маршрута для sendServiceStatus
	checkerRouter.HandleFunc("/sendServiceStatus", servServer.SendServiceStatus).Methods("POST")

	router.PathPrefix("/checker").Handler(checkerRouter)

	return servServer
}

// Start запускает HTTP сервер
func (s *ServiceServer) Start() error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		s.logger.Info("starting Rest server", zap.String("port", s.server.Addr))
		return s.server.ListenAndServe()
	})

	return eg.Wait()
}

// Stop останавливает HTTP сервер
func (s *ServiceServer) Stop() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := s.server.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("failed to shutdown rest server: %v", err)
	}

	return nil
}

func (s *ServiceServer) SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var request subscribeCheckerRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Error("failed to parse request", zap.Error(err))

		respondErr := jsonUtils.RespondWith400(w, "bad request body")
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
		}
	}

	id, err := uuid.Parse(request.CheckerUUID)
	if err != nil {
		s.logger.Error("failed to parse checker UUID", zap.Error(err))
		respondErr := jsonUtils.RespondWith400(w, "bad request body")
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
		}
	}

	err = s.CheckerService.SetCheckerAddress(id, request.Ip, request.Port)

	if err != nil {
		s.logger.Error("failed to set checker address", zap.Error(err))
		respondErr := jsonUtils.RespondWith500(w)
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
		}
	}
}

func (s *ServiceServer) SendServiceStatus(w http.ResponseWriter, r *http.Request) {
	var request sendServiceStatusRequest

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Error("failed to parse request", zap.Error(err))

		respondErr := jsonUtils.RespondWith400(w, "bad request body")
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
		}
	}

	reqUUid, err := uuid.Parse(request.RequestUUID)
	if err != nil {
		s.logger.Error("failed to parse request", zap.Error(err))
		if err := jsonUtils.RespondWith400(w, "bad request body"); err != nil {
			s.logger.Error("failed to write response", zap.Error(err))
		}
	}

	if err := s.VulnServiceService.UpdateByRequstUUID(reqUUid, request.StatusCode, request.Message, request.LastCheck); err != nil {
		s.logger.Error("failed to update vulnService", zap.Error(err))
		if _, err := w.Write([]byte("failed to update vulnService by request")); err != nil {
			s.logger.Error("failed to write response", zap.Error(err))
		}
		err := jsonUtils.RespondWith500(w)
		if err != nil {
			s.logger.Error("failed to respond", zap.Error(err))
		}
	}
}
