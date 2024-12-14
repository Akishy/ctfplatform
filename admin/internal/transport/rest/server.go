package rest

import (
	"context"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/config"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/middlewares"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/service"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/transport/rest/handlers"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Repository interface {
	service.UserRepo
	Close() error
}

// Server структура сервера
type Server struct {
	handler *echo.Echo
	server  *http.Server
	context context.Context
}

// NewServer создает новый экземпляр сервера
func NewServer(ctx context.Context, config *config.Config, repo Repository) *Server {
	// Создаем новый экземпляр Echo
	e := echo.New()

	// Middleware
	l := logger.GetLoggerFromCtx(ctx)
	e.Use(middlewares.LoggerMiddleware(l))
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Создаем экземпляр сервера
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", config.AppPort),
		Handler:      e,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	jwtService := service.NewJwtService(config.SecretKey)
	userService := service.NewUserService(repo, jwtService) // даун каст repo в подинтефейс (service.UserRepo, он в композиции основного)
	teamService := service.NewTeamService()

	server := &Server{
		handler: e,
		server:  srv,
		context: ctx,
	}

	// Настройка маршрутов перед стартом
	services := handlers.NewServices(userService, teamService)
	handlers.RegisterRoutes(ctx, e, services)

	return server
}

// Start запускает сервер
func (s *Server) Start() error {
	eg := errgroup.Group{}

	eg.Go(func() error {
		logger.GetLoggerFromCtx(s.context).Info(s.context, "starting Rest server", zap.String("port", s.server.Addr))
		return s.server.ListenAndServe()
	})

	return eg.Wait()
}

// GracefulStop корректная остановка сервера
func (s *Server) GracefulStop() error {
	// Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(s.context, 10*time.Second)
	defer cancel()

	l := logger.GetLoggerFromCtx(ctx)

	l.Info(ctx, "stopping rest server...")

	// Остановка сервера
	if err := s.server.Shutdown(ctx); err != nil {
		err = fmt.Errorf("failed to shutdown rest server: %w", err)
		return err
	}

	l.Info(ctx, "rest server stopped")
	return nil
}
