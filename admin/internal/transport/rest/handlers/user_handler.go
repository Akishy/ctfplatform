package handlers

import (
	"context"
	"errors"
	"github.com/labstack/echo/v4"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
	serviceErrors "gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/errors"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

type UserService interface {
	RegistrateUser(ctx context.Context, user *entities.User) error
	LoginUser(ctx context.Context, user *entities.User) (string, error)
}

type UserEndpoints struct {
	context context.Context
	service UserService
}

func NewUserEndpoints(ctx context.Context, service UserService) *UserEndpoints {
	return &UserEndpoints{
		context: ctx,
		service: service,
	}
}

// registrationHandler обрабатывает регистрацию пользователя
func (e *UserEndpoints) registrationHandler(c echo.Context) error {
	l := logger.GetLoggerFromCtx(e.context)

	// Создаем структуру для привязки входящих данных
	var request userRegistrationRequest

	// Привязываем JSON из запроса к структуре пользователя
	if err := c.Bind(&request); err != nil {
		l.Error(e.context, "failed to bind user", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user data")
	}

	// Валидация данных (можно добавить дополнительную проверку)
	if request.Username == "" || request.Password == "" {
		l.Error(e.context, "username or password is empty")
		return echo.NewHTTPError(http.StatusBadRequest, "username and password are required")
	}

	//перемапивание в сущность
	user := &entities.User{Username: request.Username, Password: request.Password}

	// Вызываем метод регистрации сервиса
	err := e.service.RegistrateUser(e.context, user)
	if err != nil {
		// Обработка ошибки регистрации
		if errors.Is(err, serviceErrors.ErrUserExists) {
			l.Error(e.context, err.Error())
			return echo.NewHTTPError(http.StatusConflict, err.Error())
		}

		l.Error(e.context, "failed to create user", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create user")
	}

	// Успешная регистрация
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

// loginHandler обрабатывает аутентификацию пользователя
func (e *UserEndpoints) loginHandler(c echo.Context) error {
	l := logger.GetLoggerFromCtx(e.context)

	// Создаем структуру для привязки входящих данных
	var request userLoginRequest

	// Привязываем JSON из запроса к структуре пользователя
	if err := c.Bind(&request); err != nil {
		l.Error(e.context, "failed to bind user", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid user data")
	}

	// Валидация данных (можно добавить дополнительную проверку)
	if request.Username == "" || request.Password == "" {
		l.Error(e.context, "username or password is empty")
		return echo.NewHTTPError(http.StatusBadRequest, "username and password are required")
	}

	//перемапивание в сущность
	user := &entities.User{Username: request.Username, Password: request.Password}

	// Вызываем метод входа сервиса
	token, err := e.service.LoginUser(e.context, user)
	if err != nil {
		if errors.Is(err, serviceErrors.ErrUnregisteredUser) || errors.Is(err, serviceErrors.ErrInvalidPassword) {
			l.Error(e.context, err.Error())
			return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
		}

		// Обработка ошибки аутентификации
		l.Error(e.context, "failed to authenticate user", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to authenticate user")
	}

	// Успешная аутентификация - возвращаем JWT-токен
	return c.JSON(http.StatusOK, userLoginResponse{JwtToken: token})
}
