package http

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/checker/internal/jsonUtils"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type ServiceServer struct {
	server *http.Server
	logger *zap.Logger
}

// New создает новый экземпляр ServiceServer
func New(port int, logger *zap.Logger) *ServiceServer {

	router := mux.NewRouter()
	checkerRouter := router.PathPrefix("/checker").Subrouter()

	// Регистрация маршрута для sendServiceStatus
	checkerRouter.HandleFunc("/subscribe", func(w http.ResponseWriter, r *http.Request) {
		var request subscribeCheckerRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("failed to parse request", zap.Error(err))

			respondErr := jsonUtils.RespondWith400(w, "bad request body")
			if respondErr != nil {
				logger.Error("failed to write response", zap.Error(err))
			}
		}

	}).Methods("POST")

	// Регистрация маршрута для sendServiceStatus
	checkerRouter.HandleFunc("/sendServiceStatus", func(w http.ResponseWriter, r *http.Request) {
		var request sendServiceStatusRequest

		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			logger.Error("failed to parse request", zap.Error(err))

			respondErr := jsonUtils.RespondWith400(w, "bad request body")
			if respondErr != nil {
				logger.Error("failed to write response", zap.Error(err))
			}
		}

	}).Methods("POST")

	router.PathPrefix("/checker").Handler(checkerRouter)

	return &ServiceServer{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: router,
		},
		logger: logger,
	}
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
