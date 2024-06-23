package handlers

import (
	"fmt"

	"github.com/a-h/templ"
	"github.com/donaldgifford/tbccgolf/services"
	"github.com/donaldgifford/tbccgolf/views/score"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type ScoreService interface {
	// Create(players []services.Player, match services.Match) error
	Create(playerID uint, matchID uint) error
	GetAll() ([]*services.Score, error)
	Get(scoreID int) (services.Score, error)
}

func NewScoreHandler(ss ScoreService) *ScoreHandler {
	return &ScoreHandler{ss}
}

type ScoreHandler struct {
	ScoreService ScoreService
}

func (sh *ScoreHandler) CreateScore(c echo.Context) error {
	return nil
}

func (sh *ScoreHandler) GetScore(c echo.Context) error {
	return nil
}

func (sh *ScoreHandler) GetScores(c echo.Context) error {
	scores, err := sh.ScoreService.GetAll()
	if err != nil {
		return err
	}

	fmt.Println(scores)

	titlePage := fmt.Sprintf(
		"| %s",
		cases.Title(language.English).String(c.Get(username_key).(string)),
	)

	return sh.View(c, score.ShowIndex(
		titlePage,
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		score.Show(scores),
	))
}

func (sh *ScoreHandler) ScoreDetails(c echo.Context) error {
	return nil
}

func (sh *ScoreHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
