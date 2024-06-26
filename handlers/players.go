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

	"github.com/donaldgifford/tbccgolf/views/player"
	"github.com/labstack/echo/v4"
)

type PlayerService interface {
	GetAll() ([]*services.Player, error)
	GetPlayerById(id int) (services.Player, error)
	// Update(player services.Player) error
	Create(p services.Player) error
}

func NewPlayerHandler(ps PlayerService) *PlayerHandler {
	return &PlayerHandler{ps}
}

type PlayerHandler struct {
	PlayerService PlayerService
}

func (ph *PlayerHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (ph *PlayerHandler) ShowPlayers(c echo.Context) error {
	players, err := ph.PlayerService.GetAll()
	if err != nil {
		return err
	}

	titlePage := fmt.Sprintf(
		"| Players",
	)

	return ph.View(c, player.ShowIndex(
		titlePage,
		player.Show(players),
	))
}

func (ph *PlayerHandler) CreatePlayer(c echo.Context) error {
	if c.Request().Method == "POST" {
		var player services.Player
		hcp, err := strconv.Atoi(c.FormValue("handicap"))
		if err != nil {
			loggy.Loggy().Error("Couldnt convert handicap, set to 0")
			player.Email = c.FormValue("email")
			player.Name = c.FormValue("name")
			player.Handicap = 0

		} else {

			player := services.Player{
				Email:    c.FormValue("email"),
				Name:     c.FormValue("name"),
				Handicap: hcp,
			}

			err := ph.PlayerService.Create(player)
			if err != nil {
				if strings.Contains(err.Error(), "UNIQUE constraint failed") {
					err = errors.New("this email is already registered")

					return c.Redirect(http.StatusSeeOther, "/players")

				}

				return echo.NewHTTPError(
					echo.ErrInternalServerError.Code,
					fmt.Sprintf("Something went wrong: %s", err))
			}

			return c.Redirect(http.StatusSeeOther, "/players")
		}
	}

	return ph.View(c, player.ShowIndex(
		"| Create",
		player.Create(),
	))
}

func (ph *PlayerHandler) ShowPlayerById(c echo.Context) error {
	idParam, _ := strconv.Atoi(c.Param("id"))

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}

	pdata, err := ph.PlayerService.GetPlayerById(idParam)
	if err != nil {
		if strings.Contains(err.Error(), "no rows in result set") {
			return echo.NewHTTPError(http.StatusNotFound, err)
		}
		return err
	}

	return ph.View(c, player.DetailsIndex(
		"",
		player.Details(tz, pdata),
	))
}

func (ph *PlayerHandler) APIGetAll(c echo.Context) error {
	players, err := ph.PlayerService.GetAll()
	if err != nil {
		return err
	}

	response := map[string]interface{}{
		"data": players,
	}

	return c.JSON(http.StatusOK, response)
}

func (ph *PlayerHandler) APIGetAllHTML(c echo.Context) error {
	players, err := ph.PlayerService.GetAll()
	if err != nil {
		return err
	}

	return ph.View(c, player.ShowHTML(players))
}
