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

func RegisterRoutes(ctx context.Context, handler *echo.Echo, authMiddleware echo.MiddlewareFunc, service *Services) {
	// User
	userApi := handler.Group("/user")
	userEndpoints := NewUserEndpoints(ctx, service)

	userApi.POST("/signup", userEndpoints.RegistrationHandler)
	userApi.POST("/login", userEndpoints.loginHandler)

	// Team - don't finished
	//teamApi := api.Group("/teams")
	//teamApi.Use(authMiddleware)
	//teamEndpoints := NewTeamEndpoints(ctx, service)
	//
	//teamApi.POST("", teamEndpoints.registrationHandler)
	//teamApi.GET("", teamEndpoints.getTeamsHandler)
	//teamApi.GET("/:team_id/invite", teamEndpoints.getInviteLinkHandler)
	//teamApi.POST("/:team_id/join/:invite_key", teamEndpoints.addMemberHandler)
	//teamApi.POST("/:team_id/leave", teamEndpoints.leaveTeamHandler)
	//teamApi.POST("/:team_id/kick", teamEndpoints.kickMemberHandler)
	//teamApi.DELETE("/:team_id", teamEndpoints.deleteTeamHandler)
	//teamApi.PATCH("/:team_id", teamEndpoints.editTeamHandler)
}
