package handlers

import (
	"context"
	"github.com/labstack/echo/v4"
)

type Services struct {
	UserService
	TeamService
}

func NewServices(userService UserService, teamService TeamService) *Services {
	return &Services{
		UserService: userService,
		TeamService: teamService,
	}
}

func RegisterRoutes(ctx context.Context, handler *echo.Echo, service *Services) {
	api := handler.Group("/api")

	userEndpoints := NewUserEndpoints(ctx, service)

	api.POST("/register", userEndpoints.registrationHandler)
	api.POST("/login", userEndpoints.loginHandler)
}
