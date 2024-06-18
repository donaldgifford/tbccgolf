package main

import (
	"github.com/donaldgifford/tbccgolf/db"
	"github.com/donaldgifford/tbccgolf/handlers"
	"github.com/donaldgifford/tbccgolf/services"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

const (
	SECRET_KEY = "secret"
)

func main() {
	e := echo.New()

	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	e.Static("/", "assets")

	e.Use(session.Middleware(sessions.NewCookieStore([]byte(SECRET_KEY))))

	db.Init()
	gorm := db.DB()

	dbGorm, err := gorm.DB()
	if err != nil {
		panic(err)
	}

	db.AutoMigrate()
	dbGorm.Ping()

	ps := services.NewServicesPlayer(services.Player{}, gorm)
	ms := services.NewServicesMatch(services.Match{}, gorm)
	// se := services.NewServicesEvent(services.Event{}, gorm)
	ah := handlers.NewAuthHandler(ps)

	// s := handlers.SSEHandler(se)
	p := handlers.NewPlayerHandler(ps)
	m := handlers.NewMatchHandler(ms)

	handlers.SetupRoutes(e, p, m, ah)

	e.Logger.Fatal(e.Start(":8080"))
}
