package repository

import (
	"context"
	"database/sql"
	"fmt"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/entities"
)

func (r *Repository) IsUserExistsByUsername(ctx context.Context, username string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM public.users WHERE username = $1)`

	var exists bool
	err := r.db.GetContext(ctx, &exists, query, username)

	if err != nil {
		return false, fmt.Errorf("failed to check user existence: %w", err)
	}

	return exists, nil
}

func (r *Repository) CreateUser(ctx context.Context, user *entities.User) error {
	query := `INSERT INTO public.users (username, password) VALUES ($1, $2)`

	_, err := r.db.ExecContext(ctx, query, user.Username, user.Password)
	if err != nil {
		err = fmt.Errorf("failed to create user: %w", err)
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
