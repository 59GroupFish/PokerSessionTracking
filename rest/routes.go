package rest

import (
	"github.com/labstack/echo/v4"
)

func FrontEndRoutes(e *echo.Echo) {
	// Serve the frontend files
	e.Static("/static", "./client/static")

	e.GET("/", GetHome)
	e.GET("/view-games", ViewGames)
	e.GET("/manage-players/:gameId", ManagePlayers)
}

func BackendRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/games", GetGamesHandler)
	api.PUT("/game", CreateGameHandler)
	api.PUT("/player", AddPlayerHandler)
	api.GET("/simpledebt", GetSimpleDebt)
}
