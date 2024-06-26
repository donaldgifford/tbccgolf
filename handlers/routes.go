package handlers

import "github.com/labstack/echo/v4"

func SetupRoutes(
	e *echo.Echo,
	p *PlayerHandler,
	m *MatchHandler,
	s *ScoreHandler,
) {
	e.GET("/", HomeHandler)
	playerGroup := e.Group("/players")
	playerGroup.GET("", p.ShowPlayers)
	playerGroup.GET("/create", p.CreatePlayer)
	playerGroup.POST("/create", p.CreatePlayer)
	playerGroup.GET("/details/:id", p.ShowPlayerById)

	matchGroup := e.Group("/matches")
	matchGroup.GET("", m.ShowMatches)
	matchGroup.GET("/create", m.CreateMatch)
	matchGroup.POST("/create", m.CreateMatch)
	// matchGroup.POST("/create", m.Create)
	matchGroup.GET("/details/:id", m.MatchDetails)

	matchGroup.GET("/scoring/:id", m.MatchScoringHome)

	// json api
	jsonApiGroup := e.Group("/api/json")
	jsonApiGroup.GET("/players", p.APIGetAll)

	// html api
	htmlApiGroup := e.Group("/api/html")
	htmlApiGroup.GET("/players", p.APIGetAllHTML)

	scoreGroup := e.Group("scores")
	scoreGroup.GET("", s.GetScores)
	scoreGroup.GET("/:id", s.Get)
	scoreGroup.GET("/details/:id", s.GetDetails)
	scoreGroup.POST("/create/:matchID", s.CreateScore)

	// protectedGroup.GET("/create", p.Create)
	// protectedGroup.POST("/create", p.Create)
	// protectedGroup.GET("/details/:id", p.ShowPlayerById)
	// protectedGroup.GET("/edit/:id", p.Update)
	// protectedGroup.POST("/edit/:id", p.updatePlayerHandler)
	// protectedGroup.POST("/logout", p.logoutHandler)
	// matchGroup := e.Group("/matches")
	// matchGroup.GET("", m.ShowMatches)
	// matchGroup.GET("/create", m.CreateMatch)
	// matchGroup.POST("/create", m.CreateMatch)
	// matchGroup.GET("/details/:id", m.MatchDetails)
	// scoreGroup := e.Group("/scores")
	// scoreGroup.GET("", s.GetScores)
	// scoreGroup.POST("/create", s.CreateScore)
	// sse := e.Group("/sse", ah.authMiddleware)
	// sse.GET("", s.ShowEvents)
}
