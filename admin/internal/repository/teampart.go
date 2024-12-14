package repository

import (
	"context"
	"fmt"
)

func (r *Repository) CreateTeam(ctx context.Context, name string) (bool, error) {

}

func (r *Repository) IsTeamExistsByName(ctx context.Context, name string) (bool, error) {
	query := `SELECT EXISTS(SELECT 1 FROM public.teams WHERE name = $1)`

	var exists bool
	err := r.db.GetContext(ctx, &exists, query, name)

	if err != nil {
		return false, fmt.Errorf("failed to check team existence: %w", err)
	}

	return exists, nil
}
