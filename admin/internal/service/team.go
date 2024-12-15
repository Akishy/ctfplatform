package service

import (
	"context"
)

type TeamRepo interface {
	//IsUserExistsById(ctx context.Context, id int64) (bool, error)
	//IsTeamExistsByName(ctx context.Context, name string) (bool, error)
	//TransactCreateTeam(ctx context.Context, name string, ownerId int64) error
}

type TeamService struct {
	repo TeamRepo
}

func NewTeamService(repo TeamRepo) *TeamService {
	return &TeamService{
		repo: repo,
	}
}

func (s *TeamService) RegistrateTeam(ctx context.Context, teamName string, ownerId int64) error {
	//isExists, err := s.repo.IsUserExistsById(ctx, ownerId)
	//if err != nil {
	//	return fmt.Errorf("failed to find user: %w", err)
	//}
	//
	//if isExists {
	//	return errors.ErrUnregisteredUser
	//}
	//
	//isExists, err = s.repo.IsTeamExistsByName(ctx, teamName)
	//if err != nil {
	//	return fmt.Errorf("failed to find team: %w", err)
	//}
	//
	//if isExists {
	//	return errors.ErrTeamExists
	//}
	//
	//err = s.repo.TransactCreateTeam(ctx, teamName, ownerId)
	//if err != nil {
	//	return fmt.Errorf("failed to create team: %w", err)
	//}

	return nil
}
