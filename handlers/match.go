package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/donaldgifford/tbccgolf/services"

	"github.com/donaldgifford/tbccgolf/views/match"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type MatchService interface {
	GetMatches() ([]*services.Match, error)
	GetMatch(matchID int) (services.Match, error)
	CreateMatch(m services.Match) error
	CompleteMatch(matchID int) error
	ListPlayers() ([]services.PlayerList, error)
	GetPlayer(name string) (*services.Player, error)
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

	fmt.Println(matches)
	// for _, m := range matches {
	// 	fmt.Printf("Netscore: %v\n", m.NetScore)
	// 	for _, p := range m.Players {
	// 		fmt.Println("Players=======")
	// 		fmt.Printf("Player: %s\n", p.Username)
	// 	}
	// }

	titlePage := fmt.Sprintf(
		"| %s",
		cases.Title(language.English).String(c.Get(username_key).(string)),
	)

	return mh.View(c, match.ShowIndex(
		titlePage,
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		match.Show(matches),
	))
}

type MatchCreateForm struct {
	Player1     string `form:"Player1"`
	Player2     string `form:"Player2"`
	NetScore    string `form:"scoring-type"`
	ScoringType string `form:"scoring-type"`
	Length      string `form:"Length"`
}

func (mh *MatchHandler) CreateMatch(c echo.Context) error {
	// we want to return a list of the players to populate the drop down to
	// select whos in the match.
	isError = false

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}
	cdata, err := mh.MatchService.ListPlayers()
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
	fmt.Println(cdata)

	if c.Request().Method == "POST" {
		var newMatch MatchCreateForm
		// fmt.Println(c.Request().Body)

		err := c.Bind(&newMatch)
		if err != nil {
			return c.String(http.StatusBadRequest, "bad request")
		}

		var players []*services.Player

		player1, err := mh.MatchService.GetPlayer(newMatch.Player1)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		players = append(players, player1)

		player2, err := mh.MatchService.GetPlayer(newMatch.Player2)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		players = append(players, player2)

		// conver to bools
		l, err := strconv.Atoi(newMatch.Length)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		fmt.Printf("length: %d\n", l)

		newMatchService := services.Match{
			ScoringType: newMatch.ScoringType,
			Length:      l,
			Players:     players,
		}

		// err := ah.PlayerService.CreatePlayer(player)
		err = mh.MatchService.CreateMatch(newMatchService)
		if err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}

		// fmt.Println(newMatch)

		setFlashmessages(c, "success", "Match created successfully!!")

		return c.Redirect(http.StatusSeeOther, "/matches")
		// return c.Redirect(http.StatusSeeOther, "/matches/details/"+strconv.Itoa(int(newMatch.ID)))

	}
	return mh.View(c, match.ShowIndex(
		"| Create Match",
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		match.NewMatch(cdata, tz),
	))
}
