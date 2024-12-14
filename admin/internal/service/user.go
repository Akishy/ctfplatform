package service

import (
	"context"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/errors"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entities.User) error
	FindUserByUsername(ctx context.Context, username string) (*entities.User, error)
}

type UserService struct {
	repo       UserRepo
	jwtService *JwtService
}

func NewUserService(repo UserRepo, jwtService *JwtService) *UserService {
	return &UserService{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *UserService) RegistrationUser(ctx context.Context, user *entities.User) error {
	err := user.SetHashedPassword()
	if err != nil {
		return fmt.Errorf("cannot set hashed password: %w", err)
	}

	if err = s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) LoginUser(ctx context.Context, user *entities.User) (string, error) {
	// Проверка пользователя в базе данных
	existingUser, err := s.repo.FindUserByUsername(ctx, user.Username)
	if err != nil {
		return "", errors.ErrUnknownUser
	}

	// Проверка пароля
	if !existingUser.ComparePasswords(user.Password) {
		return "", errors.ErrInvalidPassword
	}

	// Генерация JWT-токена
	token, err := s.jwtService.generateJWTToken(existingUser)
	if err != nil {
		return "", err
	}

	// Возвращаем пользователя с токеном в поле Username
	return token, nil
}
