package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/donaldgifford/tbccgolf/loggy"
	"github.com/donaldgifford/tbccgolf/services"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	"github.com/donaldgifford/tbccgolf/views/match"
)

type MatchService interface {
	GetMatches() ([]*services.Match, error)
	DbQuery() *gorm.DB
	GetMatch(matchID int) (services.Match, error)
	Create(m services.Match) error
	// CompleteMatch(matchID int) error
	// ListPlayers() ([]services.PlayerList, error)
	// GetPlayer(name string) (services.Player, error)
}

func NewMatchHandler(ms MatchService) *MatchHandler {
	return &MatchHandler{ms}
}

type MatchHandler struct {
	MatchService MatchService
}

func (mh *MatchHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (mh *MatchHandler) ShowMatches(c echo.Context) error {
	matches, err := mh.MatchService.GetMatches()
	if err != nil {
		return err
	}

	return mh.View(c, match.ShowIndex(
		"",
		match.Show(matches),
	))
}

type MatchCreateForm struct {
	Player1  string `form:"Player1"`
	Player2  string `form:"Player2"`
	Player3  string `form:"Player3"`
	Player4  string `form:"Player4"`
	GameType string `form:"GameType"`
	Scoring  string `form:"Scoring"`
	Holes    string `form:"Holes"`
	Title    string `form:"Title"`
}

func (mcf *MatchCreateForm) playerIDList() ([]int, error) {
	var returnList []int

	var stringListPlayer []string

	stringListPlayer = append(stringListPlayer, mcf.Player1)
	stringListPlayer = append(stringListPlayer, mcf.Player2)
	stringListPlayer = append(stringListPlayer, mcf.Player3)
	stringListPlayer = append(stringListPlayer, mcf.Player4)

	for _, p := range stringListPlayer {
		if p != "" {
			rp, err := strconv.Atoi(p)
			if err != nil {
				loggy.Loggy().Error(err)
				return []int{}, err

			}
			returnList = append(returnList, rp)
		} else if p == "" {
			loggy.Loggy().Info("Player id empty, skipping")
			continue
		} else {

			err := errors.New("Player id error")
			loggy.Loggy().Error(err)
			return []int{}, err
		}
	}

	return returnList, nil
}

func (mh *MatchHandler) CreateMatch(c echo.Context) error {
	// we want to return a list of the players to populate the drop down to
	// select whos in the match.

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	if c.Request().Method == "POST" {
		var newMatch MatchCreateForm

		err := c.Bind(&newMatch)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		playerIDs, err := newMatch.playerIDList()
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		var players []services.Player
		if err := mh.MatchService.DbQuery().Find(&players, playerIDs); err.Error != nil {
			loggy.Loggy().Error(err.Error)

			return c.String(http.StatusBadRequest, err.Error.Error())
		}

		newMatchService := services.Match{
			Players:  players,
			Title:    newMatch.Title,
			GameType: newMatch.GameType,
			Holes:    newMatch.Holes,
			Scoring:  newMatch.Scoring,
		}
		err = mh.MatchService.Create(newMatchService)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		return c.Redirect(http.StatusSeeOther, "/matches")
	}
	return mh.View(c, match.CreateIndex(
		"| Create Match",
		match.NewMatch(tz),
	))
}

func (mh *MatchHandler) MatchDetails(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	fmt.Println(idParam)

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}
	matchData, err := mh.MatchService.GetMatch(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}
	fmt.Println(matchData)
	fmt.Println(matchData.Players)
	fmt.Println(matchData.Scores)
	for _, p := range matchData.Players {
		fmt.Println(p.Name)
	}

	return mh.View(c, match.DetailsIndex(
		"| Match Details",
		match.Details(tz, matchData),
	))
}

func (mh *MatchHandler) MatchScoringHome(c echo.Context) error {
	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	fmt.Println(idParam)
	//
	// tz := ""
	// if len(c.Request().Header["X-Timezone"]) != 0 {
	// 	tz = c.Request().Header["X-Timezone"][0]
	// }
	matchData, err := mh.MatchService.GetMatch(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return echo.NewHTTPError(
			echo.ErrInternalServerError.Code,
			fmt.Sprintf(
				"something went wrong: %s",
				err,
			))
	}
	fmt.Println(matchData)
	fmt.Println(matchData.Players)
	fmt.Println(matchData.Scores)
	for _, p := range matchData.Players {
		fmt.Println(p.Name)
	}

	return mh.View(c, match.MatchScoringIndex(
		"| Match Scoring",
		match.MatchScoringHome(matchData),
	))
}
