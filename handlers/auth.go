package handlers

import (
	"github.com/a-h/templ"

	"github.com/donaldgifford/tbccgolf/views/auth_views"
	"github.com/labstack/echo/v4"
)

func HomeHandler(c echo.Context) error {
	return renderView(c, auth_views.HomeIndex(
		"| Home",
		auth_views.Home(),
	))
}

func renderView(c echo.Context, cmp templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)

	return cmp.Render(c.Request().Context(), c.Response().Writer)
}
