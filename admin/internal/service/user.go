package service

import (
	"context"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/errors"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/jwtutils"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *entities.User) error
	FindUserByUsername(ctx context.Context, username string) (*entities.User, error)
	IsUserExistsByUsername(ctx context.Context, username string) (bool, error)
}

type UserService struct {
	repo       UserRepo
	jwtService *jwtutils.JwtUtils
}

func NewUserService(repo UserRepo, jwtService *jwtutils.JwtUtils) *UserService {
	return &UserService{
		repo:       repo,
		jwtService: jwtService,
	}
}

func (s *UserService) RegistrateUser(ctx context.Context, user *entities.User) error {
	isExists, err := s.repo.IsUserExistsByUsername(ctx, user.Username)
	if err != nil {
		return fmt.Errorf("failed to find user: %w", err)
	}

	if isExists {
		return errors.ErrUserExists
	}

	err = user.SetHashedPassword()
	if err != nil {
		return fmt.Errorf("cannot set hashed password: %w", err)
	}

	if err = s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (s *UserService) LoginUser(ctx context.Context, user *entities.User) (string, error) {
	isExists, err := s.repo.IsUserExistsByUsername(ctx, user.Username)
	if err != nil {
		return "", fmt.Errorf("failed to find user: %w", err)
	}

	if !isExists {
		return "", errors.ErrUnregisteredUser
	}

	// Проверка пользователя в базе данных
	existingUser, err := s.repo.FindUserByUsername(ctx, user.Username)
	if err != nil {
		return "", err
	}

	// Проверка пароля
	if !existingUser.ComparePasswords(user.Password) {
		return "", errors.ErrInvalidPassword
	}

	// Генерация JWT-токена
	token, err := s.jwtService.GenerateJWTToken(existingUser.ID, existingUser.Username)
	if err != nil {
		return "", err
	}

	// Возвращаем пользователя с токеном в поле Username
	return token, nil
}
