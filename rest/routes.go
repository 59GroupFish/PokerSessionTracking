package rest

import (
	"github.com/labstack/echo/v4"
)

func FrontEndRoutes(e *echo.Echo) {
	// Serve the frontend files
	e.Static("/static", "./client/static")

	e.GET("/", GetHome)
	e.GET("/view-games", ViewGames)
	// e.GET("/manage-game", ManageGame)
}

func BackendRoutes(e *echo.Echo) {
	api := e.Group("/api")
	api.PUT("/create-game", CreateGameHandler)
	api.GET("/get-games", GetGamesHandler)
	api.GET("/simpledebt", GetSimpleDebt)
}
