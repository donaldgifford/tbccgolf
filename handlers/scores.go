package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/a-h/templ"
	"github.com/donaldgifford/tbccgolf/services"
	"github.com/donaldgifford/tbccgolf/views/score_views"
	"github.com/labstack/echo/v4"
)

type ScoreService interface {
	// Create(players []services.Player, match services.Match) error
	Create(matchID uint) error
	// GetAll() ([]*services.Score, error)
	Get(scoreID int) (services.Score, error)
	GetScores(playerID uint, matchID uint) (services.Score, error)
}

func NewScoreHandler(ss ScoreService) *ScoreHandler {
	return &ScoreHandler{ss}
}

type ScoreHandler struct {
	ScoreService ScoreService
}

func (sh *ScoreHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (sh *ScoreHandler) CreateScore(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("matchID"))

	err := sh.ScoreService.Create(uint(idParam))
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	return nil
}

func (sh *ScoreHandler) Get(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	score, err := sh.ScoreService.Get(idParam)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	fmt.Println(score)

	return nil
}

func (sh *ScoreHandler) GetDetails(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	score, err := sh.ScoreService.Get(idParam)
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	fmt.Println(score)
	return sh.View(c, score_views.DetailsIndex(
		"",
		score_views.Details(tz, score),
	))
}

func (sh *ScoreHandler) GetScores(c echo.Context) error {
	playerIDparam, _ := strconv.Atoi(c.QueryParam("player_id"))
	matchIDparam, _ := strconv.Atoi(c.QueryParam("match_id"))

	score, err := sh.ScoreService.GetScores(uint(playerIDparam), uint(matchIDparam))
	if err != nil {
		fmt.Println(err.Error())
		return echo.NewHTTPError(http.StatusNotFound, err)
	}

	fmt.Println(score)

	return nil
}
