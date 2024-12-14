package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/jmoiron/sqlx"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Close() error {
	return r.db.Close()
}

func (r *Repository) CreateUser(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO public.users (username, password) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		err = fmt.Errorf("failed to create user: %w", err)

		logger.GetLoggerFromCtx(ctx).Error(ctx, err.Error())

		return err
	}

	return nil
}

func (r *Repository) FindUserByUsername(ctx context.Context, username string) (*entities.User, error) {
	query := `SELECT id, username, password FROM public.users WHERE username = $1`

	var user entities.User
	err := r.db.GetContext(ctx, &user, query, username)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user by username: %w", err)
	}

	return &user, nil
}

func (r *Repository) FindUserById(ctx context.Context, id int64) (*entities.User, error) {
	query := `SELECT id, username, password FROM public.users WHERE id = $1`

	var user entities.User
	err := r.db.GetContext(ctx, &user, query, id)

	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("user not found: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to find user by id: %w", err)
	}

	return &user, nil
}
