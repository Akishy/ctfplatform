package service

import (
	"context"
)

type TeamRepo interface {
	CreateTeam(ctx context.Context, name string) (string, error)
	deleteTeam(ctx context.Context, teamId string) error
	addMember(ctx context.Context, teamId string, userId string, isCaptain bool) error
	deleteMember(ctx context.Context, teamId string, userId string) error
	IsTeamExistsByName(ctx context.Context, name string) (bool, error)
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
