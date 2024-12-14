package service

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
	"time"
)

// JwtClaims представляет JWT-claims
type JwtClaims struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
	jwt.RegisteredClaims
}

type JwtService struct {
	secretKey []byte
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{secretKey: []byte(secretKey)}
}

// generateJWTToken создает JWT-токен для пользователя
func (s *JwtService) generateJWTToken(user *entities.User) (string, error) {

	// Устанавливаем время жизни токена
	expirationTime := time.Now().Add(24 * time.Hour) // Токен действителен 24 часа

	// Создаем claims
	claims := &JwtClaims{
		Id:   user.ID,
		Name: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			// Уникальный идентификатор токена
			ID: user.Username,
			// Время истечения токена
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			// Время выдачи токена
			IssuedAt: jwt.NewNumericDate(time.Now()),
			// Издатель токена
			Issuer: "ctf-platform",
		},
	}

	// Создаем токен с алгоритмом подписи
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// validateJWTToken проверяет валидность JWT-токена
func (s *JwtService) validateJWTToken(tokenString string) (*JwtClaims, error) {
	// Парсим и проверяем токен
	token, err := jwt.ParseWithClaims(tokenString, &JwtClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Убеждаемся, что метод подписи соответствует ожидаемому
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method")
		}
		return s.secretKey, nil
	})

	// Проверяем наличие ошибок при парсинге
	if err != nil {
		return nil, err
	}

	// Извлекаем claims, если токен валиден
	if claims, ok := token.Claims.(*JwtClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
