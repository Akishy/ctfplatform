package middlewares

import (
	"context"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/jwtutils"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
	"net/http"
	"strings"
)

type TokenValidator interface {
	ValidateJWTToken(tokenString string) (jwt.MapClaims, error)
	ExtractUserID(claims jwt.MapClaims) (int64, error)
	ExtractUsername(claims jwt.MapClaims) (string, error)
}

// JWTMiddleware middleware для проверки JWT токена
func JWTMiddleware(ctx context.Context, validator TokenValidator) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			l := logger.GetLoggerFromCtx(ctx)

			// Получаем заголовок Authorization
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Not authorized: the token is missing",
				})
			}

			// Проверяем формат заголовка (Bearer Token)
			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token format",
				})
			}

			// Извлекаем токен
			tokenString := parts[1]

			// Валидируем токен с помощью переданного валидатора
			claims, err := validator.ValidateJWTToken(tokenString)

			if err != nil {
				l.Error(ctx, err.Error())
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Failed to validate token",
				})
			}

			// Извлекаем ID и Username из токена
			userID, err := validator.ExtractUserID(claims)
			if err != nil {
				l.Info(ctx, err.Error())
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token: missing filed: user_id",
				})
			}

			username, err := validator.ExtractUsername(claims)
			if err != nil {
				l.Info(ctx, err.Error())
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token: missing filed: username",
				})
			}

			// Сохраняем данные пользователя в контексте
			c.Set(jwtutils.UserIDKey, userID)
			c.Set(jwtutils.UsernameKey, username)

			// Передаем управление следующему обработчику
			return next(c)
		}
	}
}
