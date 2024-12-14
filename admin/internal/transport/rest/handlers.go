package rest

import (
	"context"
	"github.com/labstack/echo/v4"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
	"net/http"
)

// UserService сервис
type UserService interface {
	RegistrationUser(ctx context.Context, user *entities.User) error
	LoginUser(ctx context.Context, user *entities.User) (string, error)
}

// registrationHandler обрабатывает регистрацию пользователя
func (s *Server) registrationHandler(c echo.Context) error {
	// Создаем структуру для привязки входящих данных
	var request registrationRequest

	// Привязываем JSON из запроса к структуре пользователя
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request data",
		})
	}

	// Валидация данных (можно добавить дополнительную проверку)
	if request.Username == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Username and password are required",
		})
	}

	//перемапивание в сущность
	user := &entities.User{Username: request.Username, Password: request.Password}

	// Вызываем метод регистрации сервиса
	err := s.services.RegistrationUser(s.context, user)
	if err != nil {
		// Обработка ошибки регистрации
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Registration failed",
		})
	}

	// Успешная регистрация
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "User registered successfully",
	})
}

// loginHandler обрабатывает аутентификацию пользователя
func (s *Server) loginHandler(c echo.Context) error {
	// Создаем структуру для привязки входящих данных
	var request loginRequest

	// Привязываем JSON из запроса к структуре пользователя
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request data",
		})
	}

	// Валидация данных
	if request.Username == "" || request.Password == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Username and password are required",
		})
	}

	//перемапивание в сущность
	user := &entities.User{Username: request.Username, Password: request.Password}

	// Вызываем метод входа сервиса
	token, err := s.services.LoginUser(s.context, user)
	if err != nil {
		// Обработка ошибки аутентификации
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Authentication failed",
		})
	}

	// Успешная аутентификация - возвращаем JWT-токен
	return c.JSON(http.StatusOK, loginResponse{JwtToken: token})
}

// RegisterRoutes регистрирует маршруты для обработки регистрации и входа
func (s *Server) RegisterRoutes() {
	s.handler.POST("/register", s.registrationHandler)
	s.handler.POST("/login", s.loginHandler)
}
