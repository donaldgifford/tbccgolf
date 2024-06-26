package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/donaldgifford/tbccgolf/views/error_pages"
	"github.com/labstack/echo/v4"
)

func CustomHTTPErrorHandler(err error, c echo.Context) {
	code := http.StatusInternalServerError
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
	}
	c.Logger().Error(err)

	var errorPage func() templ.Component

	switch code {
	case 401:
		errorPage = error_pages.Error401
	case 404:
		errorPage = error_pages.Error404
	case 405:
		errorPage = error_pages.Error405
	case 500:
		errorPage = error_pages.Error500
	}

	renderView(c, error_pages.ErrorIndex(
		fmt.Sprintf("| Error (%d)", code),
		errorPage(),
	))
}
