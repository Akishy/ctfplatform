package jwtutils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	UserIDKey   = "user_id"
	UsernameKey = "username"
)

type JwtUtils struct {
	secretKey []byte
	issuer    string
	duration  time.Duration
}

func NewJwtUtils(secretKey, issuer string) *JwtUtils {
	return &JwtUtils{
		secretKey: []byte(secretKey),
		issuer:    issuer,
		duration:  24 * time.Hour, // Токен действителен 24 часа по умолчанию
	}
}

// WithDuration позволяет изменить время жизни токена
func (s *JwtUtils) WithDuration(duration time.Duration) *JwtUtils {
	s.duration = duration
	return s
}

// GenerateJWTToken создает JWT-токен для пользователя
func (s *JwtUtils) GenerateJWTToken(userId int64, username string) (string, error) {
	// Устанавливаем время жизни токена
	expirationTime := time.Now().Add(s.duration)

	// Создаем claims
	claims := jwt.MapClaims{
		UserIDKey:   userId,
		UsernameKey: username,
		"exp":       expirationTime.Unix(),
		"iat":       time.Now().Unix(),
		"iss":       s.issuer,
		"jti":       username, // Уникальный идентификатор токена
	}

	// Создаем токен с алгоритмом подписи
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Подписываем токен секретным ключом
	tokenString, err := token.SignedString(s.secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return tokenString, nil
}

// ValidateJWTToken проверяет валидность JWT-токена
func (s *JwtUtils) ValidateJWTToken(tokenString string) (jwt.MapClaims, error) {
	// Парсим и проверяем токен
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Убеждаемся, что метод подписи соответствует ожидаемому
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return s.secretKey, nil
	})

	// Проверяем наличие ошибок при парсинге
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	// Извлекаем claims, если токен валиден
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

// ExtractUserID извлекает ID пользователя из claims
func (s *JwtUtils) ExtractUserID(claims jwt.MapClaims) (int64, error) {
	userID, ok := claims[UserIDKey]
	if !ok {
		return 0, fmt.Errorf("user ID not found in token claims")
	}

	// В зависимости от типа, который может вернуть jwt.MapClaims
	switch v := userID.(type) {
	case float64:
		return int64(v), nil
	case int64:
		return v, nil
	case int:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("invalid user ID type: %T", userID)
	}
}

// ExtractUsername извлекает username из claims
func (s *JwtUtils) ExtractUsername(claims jwt.MapClaims) (string, error) {
	username, ok := claims[UsernameKey].(string)
	if !ok {
		return "", fmt.Errorf("username not found in token claims")
	}
	return username, nil
}
