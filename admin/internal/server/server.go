package server

import (
	"context"
	"errors"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Server структура веб-сервера.
type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

// NewServer создает новый сервер.
func NewServer(cfg *config.Config, log *zap.Logger) *Server {
	mux := http.NewServeMux()

	// Middleware для логирования
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		// Обработчик по умолчанию
		_, err := w.Write([]byte("Hello, world!"))
		if err != nil {
			log.Error("Failed to write response", zap.Error(err))
			return
		}

		// Логируем запрос
		log.Info("Request received",
			zap.String("method", r.Method),
			zap.String("url", r.URL.String()),
			zap.String("remote_addr", r.RemoteAddr),
			zap.Duration("duration", time.Since(start)),
		)
	})

	server := &http.Server{
		Addr:    ":" + cfg.AppPort,
		Handler: mux,
	}

	return &Server{httpServer: server, logger: log}
}

// StartServer запускает сервер и добавляет хук для остановки.
func StartServer(lc fx.Lifecycle, server *Server) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				server.logger.Info("Starting server", zap.String("addr", server.httpServer.Addr))
				if err := server.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
					server.logger.Fatal("Server failed", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.logger.Info("Stopping server...")
			return server.httpServer.Shutdown(ctx)
		},
	})
}

// Module предоставляет зависимости сервера.
var Module = fx.Options(
	fx.Provide(NewServer),
	fx.Invoke(StartServer),
)
