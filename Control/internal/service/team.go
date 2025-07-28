package service

import (
	"context"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/models"
)

type TeamRepo interface {
	CreateTeamWithCaptain(ctx context.Context, teamName string, userId int64) (string, error)
	DeleteTeam(ctx context.Context, teamId string) error
	AddMember(ctx context.Context, teamId string, userId string, isCaptain bool) error
	DeleteMember(ctx context.Context, teamId string, userId string) error
	IsTeamExistsByName(ctx context.Context, name string) (bool, error)
	GetTeamMembers(ctx context.Context, teamId string) ([]models.User, error)
	GetTeams(ctx context.Context) ([]models.Team, error)
}

type TeamService struct {
	repo TeamRepo
}

func NewTeamService(repo TeamRepo) *TeamService {
	return &TeamService{
		repo: repo,
	}
}

func (s *TeamService) CreateTeam(ctx context.Context, teamName string, ownerId int64) (string, error) {
	teamId, err := s.repo.CreateTeamWithCaptain(ctx, teamName, ownerId)
	if err != nil {
		return "", err
	}

	return teamId, nil
}

func (s *TeamService) DeleteTeam(ctx context.Context, teamId string, userId string) error {

	// проверка на является пользователь капитаном которой тут нет

	err := s.repo.DeleteTeam(ctx, teamId)
	if err != nil {
		return err
	}
	return nil
}

func (s *TeamService) AddMember(ctx context.Context, teamId string, userId string) error {
	err := s.repo.AddMember(ctx, teamId, userId, false)
	if err != nil {
		return err
	}
	return nil
}

func (s *TeamService) IsTeamExistsByName(ctx context.Context, name string) (bool, error) {
	exists, err := s.repo.IsTeamExistsByName(ctx, name)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (s *TeamService) GetTeams(ctx context.Context) ([]models.Team, error) {
	teams, err := s.repo.GetTeams(ctx)
	if err != nil {
		return nil, err
	}
	return teams, nil
}

func (s *TeamService) GetTeamMembers(ctx context.Context, teamId string) ([]models.User, error) {
	members, err := s.repo.GetTeamMembers(ctx, teamId)
	if err != nil {
		return nil, err
	}
	return members, nil
}
