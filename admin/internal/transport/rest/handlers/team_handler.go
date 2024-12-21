package handlers

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/jwtutils"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/internal/models"
	"gitlab.crja72.ru/gospec/go4/ctfplatform/admin/pkg/logger"
	"go.uber.org/zap"
	"net/http"
)

type TeamService interface {
	CreateTeam(ctx context.Context, teamName string, ownerId int64) (string, error)
	DeleteTeam(ctx context.Context, teamId string, userId string) error
	AddMember(ctx context.Context, teamId string, userId string) error
	//DeleteMember(ctx context.Context, teamId string, userId string) error
	IsTeamExistsByName(ctx context.Context, name string) (bool, error)
	GetTeams(ctx context.Context) ([]models.Team, error)
	GetTeamMembers(ctx context.Context, teamId string) ([]models.User, error)
}

type TeamEndpoints struct {
	context context.Context
	service TeamService
}

func NewTeamEndpoints(ctx context.Context, service TeamService) *TeamEndpoints {
	return &TeamEndpoints{
		context: ctx,
		service: service,
	}
}

func (e *TeamEndpoints) createTeamHandler(c echo.Context) error {
	// на этом уровне не может быть ошибки кастинга
	userid, _ := c.Get(jwtutils.UsernameKey).(int64)

	l := logger.GetLoggerFromCtx(e.context)

	// Создаем структуру для привязки входящих данных
	var request teamRegistrationRequest

	// Привязываем JSON из запроса к структуре пользователя
	if err := c.Bind(&request); err != nil {
		l.Error(e.context, "failed to bind team", zap.Error(err))
		return echo.NewHTTPError(http.StatusBadRequest, "invalid team data")
	}

	// Валидация данных
	if request.TeamName == "" {
		l.Error(e.context, "team_name is empty")
		return echo.NewHTTPError(http.StatusBadRequest, "team_name are required")
	}

	// Вызываем метод регистрации сервиса
	teamId, err := e.service.CreateTeam(e.context, request.TeamName, userid)
	if err != nil {
		// обработка ошибок сервиса
		//...
		//...
		//...

		l.Error(e.context, "failed to create team", zap.Error(err))
		return echo.NewHTTPError(http.StatusInternalServerError, "failed to create team")
	}

	// Успешная регистрация
	return c.JSON(http.StatusCreated, map[string]string{
		"message": fmt.Sprintf("Team registered successfully, uuid: [%v]", teamId),
	})
}

func (e *TeamEndpoints) getTeamsHandler(c echo.Context) error {
	panic("implement me")
}

func (e *TeamEndpoints) getInviteLinkHandler(c echo.Context) error {
	c.Param("team_id")
	panic("implement me")
}

func (e *TeamEndpoints) addMemberHandler(c echo.Context) error {
	c.Param("team_id")
	c.Param("invite_key")
	panic("implement me")
}

func (e *TeamEndpoints) leaveTeamHandler(c echo.Context) error {
	c.Param("team_id")
	panic("implement me")
}

func (e *TeamEndpoints) kickMemberHandler(c echo.Context) error {
	c.Param("team_id")
	panic("implement me")
}

func (e *TeamEndpoints) deleteTeamHandler(c echo.Context) error {
	c.Param("team_id")
	panic("implement me")
}
