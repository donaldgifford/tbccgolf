package handlers

import "github.com/labstack/echo/v4"

var (
	fromProtected bool = false
	isError       bool = false
)

func SetupRoutes(
	e *echo.Echo,
	p *PlayerHandler,
	m *MatchHandler,
	// s *SSEHandler,
	ah *AuthHandler,
) {
	e.GET("/", ah.homeHandler)
	e.GET("/login", ah.loginHandler)
	e.POST("/login", ah.loginHandler)
	e.GET("/register", ah.registerHandler)
	e.POST("/register", ah.registerHandler)
	// apiGroup := e.Group("/api")
	// apiGroup.POST("/scorecard", s.APICreateScoreCard)
	protectedGroup := e.Group("/player", ah.authMiddleware)
	protectedGroup.GET("", p.HandlerShowPlayers)
	protectedGroup.GET("/create", p.CreatePlayer)
	protectedGroup.POST("/create", p.CreatePlayer)
	protectedGroup.GET("/details/:id", p.HandlerShowPlayerById)
	protectedGroup.GET("/edit/:id", p.updatePlayerHandler)
	protectedGroup.POST("/edit/:id", p.updatePlayerHandler)
	protectedGroup.POST("/logout", p.logoutHandler)
	matchGroup := e.Group("/matches", ah.authMiddleware)
	matchGroup.GET("", m.ShowMatches)
	matchGroup.GET("/create", m.CreateMatch)
	matchGroup.POST("/create", m.CreateMatch)
	// sse := e.Group("/sse", ah.authMiddleware)
	// sse.GET("", s.ShowEvents)
}
