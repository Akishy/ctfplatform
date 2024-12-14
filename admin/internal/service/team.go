package service

import "context"

type TeamService struct {
}

func NewTeamService() *TeamService {
	return &TeamService{}
}

func (s *TeamService) RegistrateTeam(ctx context.Context, teamName string) error {
	panic("implement me")
}
