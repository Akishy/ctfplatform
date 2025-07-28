package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/models"
)

func (r *Repository) CreateTeamWithCaptain(ctx context.Context, teamName string, userId int64) (string, error) {
	tx, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		return "", fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Делаем defer для отката транзакции в случае ошибки
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// Создаем команду
	inviteLink := uuid.New().String()
	teamId := uuid.New().String()
	createTeamQuery := `INSERT INTO public.teams (team_id, teamName, invite_link) VALUES ($1, $2, $3)`

	_, err = tx.ExecContext(ctx, createTeamQuery, teamId, teamName, inviteLink)
	if err != nil {
		return "", fmt.Errorf("failed to create team: %w", err)
	}

	// Добавляем капитана
	addCaptainQuery := `INSERT INTO public.teams_members (team_id, user_id, is_captain) VALUES ($1, $2, $3)`
	_, err = tx.ExecContext(ctx, addCaptainQuery, teamId, userId, true)
	if err != nil {
		return "", fmt.Errorf("failed to add captain: %w", err)
	}

	// Подтверждаем транзакцию
	if err = tx.Commit(); err != nil {
		return "", fmt.Errorf("failed to commit transaction: %w", err)
	}

	return teamId, nil
}

func (r *Repository) DeleteTeam(ctx context.Context, teamId string) error {
	query := `DELETE FROM public.teams WHERE team_id = $1`
	_, err := r.db.ExecContext(ctx, query, teamId)
	if err != nil {
		return fmt.Errorf("failed to delete team: %w", err)
	}
	return nil
}

func (r *Repository) AddMember(ctx context.Context, teamId string, userId string, isCaptain bool) error {
	query := `INSERT INTO public.teams_members (team_id, user_id, is_captain) VALUES ($1, $2, $3)`
	_, err := r.db.ExecContext(ctx, query, teamId, userId, isCaptain)
	if err != nil {
		return fmt.Errorf("failed to add member: %w", err)
	}
	return nil
}

func (r *Repository) DeleteMember(ctx context.Context, teamId string, userId string) error {
	query := `DELETE FROM public.teams_members WHERE team_id = $1 AND user_id = $2`
	_, err := r.db.ExecContext(ctx, query, teamId, userId)
	if err != nil {
		return fmt.Errorf("failed to delete member: %w", err)
	}
	return nil
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

func (r *Repository) GetTeamMembers(ctx context.Context, teamId string) ([]models.User, error) {
	query := `SELECT user_id, username FROM public.teams_members WHERE team_id = $1`
	var members []models.User
	err := r.db.SelectContext(ctx, &members, query, teamId)
	if err != nil {
		return nil, fmt.Errorf("failed to get team members: %w", err)
	}
	return members, nil
}

func (r *Repository) GetTeams(ctx context.Context) ([]models.Team, error) {
	query := `SELECT team_id, name, invite_link FROM public.teams`
	var teams []models.Team
	err := r.db.SelectContext(ctx, &teams, query)
	if err != nil {
		return nil, fmt.Errorf("failed to get teams: %w", err)
	}
	for i := range teams {
		members, err := r.GetTeamMembers(ctx, teams[i].TeamId)
		if err != nil {
			return nil, fmt.Errorf("failed to get team members: %w", err)
		}
		teams[i].Members = members
	}
	return teams, nil
}
