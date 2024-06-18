package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/a-h/templ"
	"github.com/donaldgifford/tbccgolf/services"
	"github.com/donaldgifford/tbccgolf/views/player"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type PlayerService interface {
	GetAllPlayers() ([]*services.Player, error)
	GetPlayerById(id int) (services.Player, error)
	UpdatePlayer(player services.Player) error
	CreatePlayer(p services.Player) error
}

func NewPlayerHandler(ps PlayerService) *PlayerHandler {
	return &PlayerHandler{ps}
}

type PlayerHandler struct {
	PlayerService PlayerService
}

func (ph *PlayerHandler) HandlerShowPlayers(c echo.Context) error {
	players, err := ph.PlayerService.GetAllPlayers()
	if err != nil {
		return err
	}

	titlePage := fmt.Sprintf(
		"| %s",
		cases.Title(language.English).String(c.Get(username_key).(string)),
	)
	// showView := player.Show(players)

	// si := player.ShowIndex("| Home", player.Show(players))
	// return ph.View(c, si)
	return ph.View(c, player.ShowIndex(
		// "| Home",
		titlePage,
		c.Get(username_key).(string),
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		player.Show(players),
	))
}

func (ph *PlayerHandler) HandlerShowPlayerById(c echo.Context) error {
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

	// di := player.DetailsIndex(
	// 	fmt.Sprintf("| User details %s",
	// 		cases.Title(language.English).String(pdata.Name),
	// 	),
	// 	player.Details(tz, pdata),
	// )

	// return ph.View(c, di)
	return ph.View(c, player.DetailsIndex(
		fmt.Sprintf("| User details %s",
			cases.Title(language.English).String(pdata.Name),
		),
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		player.Details(tz, pdata),
	))
}

func (ph *PlayerHandler) View(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}

func (ph *PlayerHandler) logoutHandler(c echo.Context) error {
	sess, _ := session.Get(auth_sessions_key, c)
	sess.Values = map[interface{}]interface{}{
		auth_key:     false,
		user_id_key:  "",
		username_key: "",
		tzone_key:    "",
	}
	sess.Save(c.Request(), c.Response())

	setFlashmessages(c, "success", "You have been logged out")

	fromProtected = false

	return c.Redirect(http.StatusSeeOther, "/login")
}

func (ph *PlayerHandler) updatePlayerHandler(c echo.Context) error {
	isError = false

	idParam, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	fmt.Println(idParam)

	tz := ""
	if len(c.Request().Header["X-Timezone"]) != 0 {
		tz = c.Request().Header["X-Timezone"][0]
	}
	cdata, err := ph.PlayerService.GetPlayerById(idParam)
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
		cdata.Name = c.FormValue("name")
		cdata.Email = c.FormValue("email")
		cdata.Username = c.FormValue("username")

		fmt.Println(cdata)
		err := ph.PlayerService.UpdatePlayer(cdata)
		if err != nil {
			return err
		}

		setFlashmessages(c, "success", "Task created successfully!!")

		return c.Redirect(http.StatusSeeOther, "/player/details/"+strconv.Itoa(idParam))
	}
	return ph.View(c, player.DetailsIndex(
		fmt.Sprintf("Edit Player %s", cdata.Name),
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		player.UpdatePlayer(cdata, tz),
		// player.UpdatePlayer(cdata, tz),
		// course_views.UpdateCourse(cdata, c.Get(tzone_key).(string)),
	))
}

func (ph *PlayerHandler) CreatePlayer(c echo.Context) error {
	isError = false

	if c.Request().Method == "POST" {
		player := services.Player{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Username: c.FormValue("username"),
		}

		err := ph.PlayerService.CreatePlayer(player)
		if err != nil {
			if strings.Contains(err.Error(), "UNIQUE constraint failed") {
				err = errors.New("this email is already registered")
				setFlashmessages(c, "error", fmt.Sprintf("Something went wrong: %s", err))

				return c.Redirect(http.StatusSeeOther, "/register")

			}

			return echo.NewHTTPError(
				echo.ErrInternalServerError.Code,
				fmt.Sprintf("Something went wrong: %s", err))
		}

		setFlashmessages(c, "success", "You are now registered!")
		return c.Redirect(http.StatusSeeOther, "/login")
	}

	return ph.View(c, player.ShowIndex(
		"| Create",
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		player.Create(fromProtected),
	))
}
