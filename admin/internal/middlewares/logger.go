package middlewares

import (
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func LoggerMiddleware(l logger.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// Начало замера времени
			start := time.Now()

			// Логируем входящий запрос
			l.Info(c.Request().Context(), "request started",
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.String("remote_ip", c.RealIP()),
			)

			// Выполняем следующий обработчик
			err := next(c)

			// Вычисляем длительность запроса
			duration := time.Since(start)

			// Логируем результат запроса
			logFields := []zap.Field{
				zap.String("method", c.Request().Method),
				zap.String("path", c.Path()),
				zap.Duration("duration", duration),
				zap.Int("status", c.Response().Status),
			}

			// Логируем ошибку, если она есть
			if err != nil {
				l.Error(c.Request().Context(), "request failed",
					append(logFields, zap.Error(err))...,
				)
				return err
			}

			l.Info(c.Request().Context(), "request completed", logFields...)

			return nil
		}
	}
}
