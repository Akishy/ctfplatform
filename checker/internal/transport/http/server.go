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
func New(port int, vulnServiceService *vulnServiceService.Service, checkerService *checkerService.Service, logger *zap.Logger) *ServiceServer {
	router := mux.NewRouter()
	servServer := &ServiceServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		logger:             logger,
		VulnServiceService: vulnServiceService,
		CheckerService:     checkerService,
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
	response := subscribeCheckerResponse{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Error("failed to parse request", zap.Error(err))

		respondErr := jsonUtils.RespondWith400(w, "bad request body")
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}
	s.logger.Debug("Decoded request to subscribeCheckerRequest struct", zap.Any("request", request))

	id, err := uuid.Parse(request.CheckerUUID)
	if err != nil {
		s.logger.Error("failed to parse checker UUID", zap.Error(err))
		respondErr := jsonUtils.RespondWith400(w, "bad request body")
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}
	s.logger.Debug("parsed request.CheckerUUID to uuid", zap.Any("id", id))

	if s.CheckerService.Subscribed(id) {
		response.Status = "you are already subscribed"
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(&response); err != nil {
			s.logger.Error("failed to encode response", zap.Error(err))
			respondErr := jsonUtils.RespondWith500(w)
			if respondErr != nil {
				s.logger.Error("failed to write response", zap.Error(err))
				return
			}
			return
		}
		return
	}

	err = s.CheckerService.SetCheckerAddress(id, request.Ip, request.Port)
	if err != nil {
		s.logger.Error("failed to set checker address", zap.Error(err))
		_, _ = w.Write([]byte("failed to set checker address"))
		respondErr := jsonUtils.RespondWith500(w)
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}

	response.Status = "success"
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		s.logger.Error("failed to encode response", zap.Error(err))
		respondErr := jsonUtils.RespondWith500(w)
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}
}

func (s *ServiceServer) SendServiceStatus(w http.ResponseWriter, r *http.Request) {
	var request sendServiceStatusRequest
	response := sendServiceStatusResponse{}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		s.logger.Error("failed to parse request", zap.Error(err))

		respondErr := jsonUtils.RespondWith400(w, "bad request body")
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}

	reqUUid, err := uuid.Parse(request.RequestUUID)
	if err != nil {
		s.logger.Error("failed to parse request", zap.Error(err))
		if err := jsonUtils.RespondWith400(w, "bad request body"); err != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}

	if err := s.VulnServiceService.UpdateByRequestUUID(reqUUid, request.StatusCode, request.Message, request.LastCheck); err != nil {
		s.logger.Error("failed to update vulnService", zap.Error(err))
		w.WriteHeader(http.StatusBadRequest)
		if _, err := w.Write([]byte("failed to update vulnService by request")); err != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}
	response.Accepted = "accepted"
	if err := json.NewEncoder(w).Encode(&response); err != nil {
		s.logger.Error("failed to encode response", zap.Error(err))
		respondErr := jsonUtils.RespondWith500(w)
		if respondErr != nil {
			s.logger.Error("failed to write response", zap.Error(err))
			return
		}
		return
	}
}
