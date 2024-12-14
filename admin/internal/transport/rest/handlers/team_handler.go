package handlers

import (
	"context"
)

type TeamService interface {
	RegistrateTeam(ctx context.Context, teamName string) error
}

type TeamEndpoints struct {
	context context.Context
	service UserService
}

func NewTeamEndpoints(ctx context.Context, service UserService) *TeamEndpoints {
	return &TeamEndpoints{
		context: ctx,
		service: service,
	}
}
