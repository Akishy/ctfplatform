package service

import (
	"context"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/models"
)

type TeamRepo interface {
	CreateTeam(ctx context.Context, name string) (string, error)
	deleteTeam(ctx context.Context, teamId string) error
	addMember(ctx context.Context, teamId string, userId string, isCaptain bool) error
	deleteMember(ctx context.Context, teamId string, userId string) error
	IsTeamExistsByName(ctx context.Context, name string) (bool, error)
	getTeamMembers(ctx context.Context, teamId string) ([]models.User, error)
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

func (s *TeamService) CreateTeam(ctx context.Context, teamName string, ownerId string) error {
	teamId, err := s.repo.CreateTeam(ctx, teamName)
	if err != nil {
		return err
	}
	err = s.repo.addMember(ctx, teamId, ownerId, true)
	if err != nil {
		return err
	}
	return nil
}

func (s *TeamService) DeleteTeam(ctx context.Context, teamId string, userId string) error {

	err := s.repo.deleteMember(ctx, teamId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *TeamService) AddMember(ctx context.Context, teamId string, userId string, isCaptain bool) error {
	err := s.repo.addMember(ctx, teamId, userId, isCaptain)
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
	members, err := s.repo.getTeamMembers(ctx, teamId)
	if err != nil {
		return nil, err
	}
	return members, nil
}
