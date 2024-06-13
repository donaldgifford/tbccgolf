package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/a-h/templ"
	"github.com/gorilla/sessions"

	"github.com/donaldgifford/tbccgolf/services"
	"github.com/donaldgifford/tbccgolf/views/auth_views"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	auth_sessions_key string = "authenticate-sessions"
	auth_key          string = "authenticated"
	user_id_key       string = "user_id"
	username_key      string = "username"
	tzone_key         string = "time_zone"
)

type AuthService interface {
	CreatePlayer(p services.Player) error
	CheckEmail(email string) (services.Player, error)
}

func NewAuthHandler(as AuthService) *AuthHandler {
	return &AuthHandler{PlayerService: as}
}

type AuthHandler struct {
	PlayerService AuthService
}

func (ah *AuthHandler) homeHandler(c echo.Context) error {
	homeView := auth_views.Home(fromProtected)
	isError = false

	return renderView(c, auth_views.HomeIndex(
		"| Home",
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		homeView,
	))
}

func (ah *AuthHandler) registerHandler(c echo.Context) error {
	registerView := auth_views.Register(fromProtected)
	isError = false

	if c.Request().Method == "POST" {
		player := services.Player{
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
			Username: c.FormValue("username"),
		}

		err := ah.PlayerService.CreatePlayer(player)
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

	return renderView(c, auth_views.RegisterIndex(
		"| Register",
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		registerView,
	))
}

func (ah *AuthHandler) loginHandler(c echo.Context) error {
	loginView := auth_views.Login(fromProtected)
	isError = false

	if c.Request().Method == "POST" {
		tzone := ""
		if len(c.Request().Header["X-Timezone"]) != 0 {
			tzone = c.Request().Header["X-Timezone"][0]
		}

		player, err := ah.PlayerService.CheckEmail(c.FormValue("email"))
		if err != nil {
			if strings.Contains(err.Error(), "no rows in result set") {
				setFlashmessages(c, "error", "There is no user with that email")

				return c.Redirect(http.StatusSeeOther, "/login")
			}

			return echo.NewHTTPError(
				echo.ErrInternalServerError.Code,
				fmt.Sprintf(
					"something went wrong: %s",
					err,
				))
		}

		err = bcrypt.CompareHashAndPassword(
			[]byte(player.Password),
			[]byte(c.FormValue("password")),
		)
		if err != nil {
			setFlashmessages(c, "error", "Incorrect password")

			return c.Redirect(http.StatusSeeOther, "/login")
		}

		sess, _ := session.Get(auth_sessions_key, c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   3600,
			HttpOnly: true,
		}

		sess.Values = map[interface{}]interface{}{
			auth_key:     true,
			user_id_key:  player.ID,
			username_key: player.Username,
			tzone_key:    tzone,
		}

		sess.Save(c.Request(), c.Response())

		setFlashmessages(c, "success", "You are now logged in!")

		return c.Redirect(http.StatusSeeOther, "/")
	}

	return renderView(c, auth_views.LoginIndex(
		"| Login",
		"",
		fromProtected,
		isError,
		getFlashmessages(c, "error"),
		getFlashmessages(c, "success"),
		loginView,
	))
}

func (ah *AuthHandler) authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := session.Get(auth_sessions_key, c)
		if auth, ok := sess.Values[auth_key].(bool); !ok || !auth {
			fromProtected = false
			return echo.NewHTTPError(echo.ErrUnauthorized.Code, "Unauthorized")
		}

		if userId, ok := sess.Values[user_id_key].(int); ok && userId != 0 {
			c.Set(user_id_key, userId)
		}

		if username, ok := sess.Values[username_key].(string); ok && len(username) != 0 {
			c.Set(username_key, username)
		}

		if tzone, ok := sess.Values[tzone_key].(string); ok && len(tzone) != 0 {
			c.Set(tzone_key, tzone)
		}

		fromProtected = true

		return next(c)
	}
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
